import path from 'path'
import LicensePlugin from 'webpack-license-plugin'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

export default {
  mode: 'production',
  entry: './src/index.js',
  target: 'node',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'action.cjs',
  },
  plugins: [
    new LicensePlugin(),
  ]
}
