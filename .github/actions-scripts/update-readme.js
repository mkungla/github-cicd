#!/usr/bin/env node
import { exec } from '@actions/exec'
import * as core from '@actions/core'
import * as github from '@actions/github'

async function capture(cmd, args) {
  const res = {
    stdout: '',
    stderr: '',
    code: null,
  }
  try {
    const code = await exec(cmd, args, {
      listeners: {
        stdout(data) {
          res.stdout += data.toString()
        },
        stderr(data) {
          res.stderr += data.toString()
        },
      },
    })
    res.code = code
    return res
  } catch (err) {
      const msg = `Command '${cmd}' failed with args '${args.join(' ')}': ${res.stderr}: ${err}`
      core.debug(`@actions/exec.exec() threw an error: ${msg}`)
      throw new Error(msg)
  }
}

export async function git(...args) {
  core.debug(`Executing Git: ${args.join(' ')}`)
  const userArgs = [
      '-c',
      'user.name=github-action-readme',
      '-c',
      'user.email=github@users.noreply.github.com',
      '-c',
      'http.https://github.com/.extraheader=', // This config is necessary to support actions/checkout@v2 (#9)
  ]
  const res = await capture('git', userArgs.concat(args))
  if (res.code !== 0) {
    throw new Error(`Command 'git ${args.join(' ')}' failed: ${JSON.stringify(res)}`)
  }
  return res.stdout
}

function getRemoteUrl() {
  const fullName = github.context.payload.repository?.full_name

  if (!fullName) {
      throw new Error(`Repository info is not available in payload: ${JSON.stringify(github.context.payload)}`)
  }

  return `https://x-access-token:${core.getInput('github-token')}@github.com/${fullName}.git`
}

export async function gitpush() {
  core.debug('Executing git push')

  const remote = getRemoteUrl()
  let args = ['push', remote, `main:main`, '--no-verify']
  if (options.length > 0) {
      args = args.concat(options)
  }

  return cmd(...args)
}

async function run() {
  await git('add', '-A')
  await git('commit', '-m', `generate and update (README.md)`)

  try {
    await gitpush()
    core.debug('git push is done')
  } catch (err) {
    core.warning('Auto-push failed')
    core.error(err)
    core.setFailed(err.message)
    process.exit(1)
  }
}

run()
