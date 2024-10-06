const path = require('path');
const glob = require('glob');

const prependDotSlash = (files) => files.map(file => `./${file}`);

module.exports = {
  entry: prependDotSlash(glob.sync('././site/home/js/*.js')),
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'static/js'),
  },
  mode: 'production',
};