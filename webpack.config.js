var webpack = require('webpack');
var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');


var APP_DIR = path.resolve(__dirname, '/javascript/src');
var OUTPUT_DIR = path.resolve(__dirname, 'out/public');

module.exports = {
  entry: './javascript/src/main.js',
  output: {
    path: OUTPUT_DIR,
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['.js', '.jsx']
  },
  module : {
    loaders: [
      { test: /\.js$/, loader: 'babel-loader', exclude: /node_modules/ },
      { test: /\.jsx$/, loader: 'babel-loader', exclude: /node_modules/ }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: 'index.ejs',
      inject: 'body',
    })
  ]
};