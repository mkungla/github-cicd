#!/usr/bin/env node

import * as core from '@actions/core'
import * as github from '@actions/github'

async function sleep(millis) {
  return new Promise(resolve => setTimeout(resolve, millis))
}

const getworkflows = (octokit) => {
  return octokit.rest.checks.listForRef({
    ref: 'main',
    owner: 'mkungla',
    repo: 'github-cicd-experiments'
  })
}
async function run() {
  const { action } = github.context
  core.info(`running ${action}`)

  const token = process.env.GITHUB_TOKEN
  // init octokit
  const octokit = github.getOctokit(token)

  const retry = 60 // times
  const interval = 5000 // 5 sec

  for (let i = 1; i <= retry; i++) {

    const {
      data: {
        total_count: running,
        check_runs: workflows,
      },
    } = await getworkflows(octokit)

    if (retry === i) {
      core.setFailed(`timeout: there is ${running} running check runs.`)
    }

    let mustWait = false
    if (running > 0) {
      for (const workflow of workflows) {
        if (workflow.check_suite.id === action) {
          continue
        }
        if (workflow.conclusion === 'failure') {
          core.setFailed(`workflow ${workflow.name} failed`)
        }
      }
      core.info(`waiting for other jobs to finish retry ${i}/${retry}`)
    }

    if (mustWait) {
      await sleep(interval)
    } else {
      break
    }
  }
}

run()
