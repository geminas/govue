module.exports = {
  entry: "./{{component}}.js",
  output: {
    path: './',
    filename: "{{component}}_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}