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

  const token = process.env.GITHUB_TOKEN
  // init octokit
  const octokit = github.getOctokit(token)

  const retry = 60 // times
  const interval = 5000 // 5 sec

  for (let i = 1; i < retry; i++) {
    const {
      data: {
        total_count: running,
        check_runs: workflows,
      },
    } = await getworkflows(octokit)

    if (running > 0) {
      for (const workflow of workflows) {
        if (workflow.conclusion === 'failure') {
          core.setFailed(`workflow ${workflow.name} failed`)
        }
      }
      core.debug(`waiting for other jobs to finish retry ${i}/${retry}`)
      await sleep(interval)
    } else {
      break
    }
  }
}

run()
