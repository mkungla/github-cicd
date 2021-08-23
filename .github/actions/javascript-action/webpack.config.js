const path = require('path')
const LicensePlugin = require('webpack-license-plugin')

module.exports = {
  mode: 'production',
  entry: './src/index.js',
  target: 'node',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'action.js',
  },
  plugins: [
    new LicensePlugin(),
  ]
}
