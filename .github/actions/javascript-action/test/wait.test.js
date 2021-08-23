import process from 'process'
import cp from 'child_process'

import path from 'path'
import chai from 'chai'
import cap from 'chai-as-promised'
import wait from '../src/wait.js'
import { fileURLToPath } from 'url'

chai.use(cap)

const { expect } = chai

const __dirname = path.dirname(fileURLToPath(import.meta.url))

describe('wait', () => {

  it('throws invalid number', async () => {
    await expect(wait('foo'))
    .to.eventually.be.rejected
    .and.be.an.instanceOf(TypeError)
    .and.have.property('message', 'milliseconds not a number');
  })

  it('wait 500 ms', async () => {
    const start = new Date()
    await wait(500)
    const end = new Date()
    var delta = Math.abs(end - start)
    expect(delta).to.be.gt(500)
  })

  // shows how the runner will run a javascript action with env / stdout protocol
  it('test runs', async () => {
    process.env['INPUT_MILLISECONDS'] = 500
    const ip = path.join(__dirname, '..', 'src', 'index.js')
    console.log(cp.execSync(`node ${ip}`, {env: process.env}).toString())
  })
})
