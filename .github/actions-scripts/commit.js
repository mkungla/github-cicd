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
      core.error(`@actions/exec.exec() threw an error: ${msg}`)
      throw new Error(msg)
  }
}

export async function git(...args) {
  core.info(`Executing Git: ${args.join(' ')}`)
  const userArgs = [
      '-c',
      'user.name=github-actions',
      '-c',
      'user.email=github-actions[bot]@users.noreply.github.com',
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
  const token = process.env.GITHUB_TOKEN
  return `https://x-access-token:${token}@github.com/${fullName}.git`
}

export async function gitpush() {
  core.info('executing git push')

  const remote = getRemoteUrl()
  let args = ['push', remote, `main:main`, '--no-verify']
  return git(...args)
}

async function run() {
  await git('add', '-A')
  // await git('commit', '-m', core.getInput('commit-message') || `auto-commit: ${process.env.GITHUB_WORKFLOW}`)
  await git('commit', '-m', process.env.MESSAGE || `auto-commit: ${process.env.GITHUB_WORKFLOW}`)


  try {
    await gitpush()
    core.info('git push is done')
  } catch (err) {
    core.notice('auto-push failed')
    core.error(err)
    core.setFailed(err.message)
    process.exit(1)
  }
}

if (!process.env.GITHUB_TOKEN) {
  core.setFailed('Error! You must have a GITHUB_TOKEN environment variable for repo access.')
  process.exit(1)
}

run()
