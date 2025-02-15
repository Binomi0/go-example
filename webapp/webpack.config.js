// webpack.config.js
const path = require("path");

module.exports = {
  mode: "development", // Development or production mode
  entry: "./src/index.tsx",
  output: {
    filename: "bundle.js", // Output file name
    path: path.resolve(__dirname, "../public/js"), // Path to the output directory
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: ["ts-loader"],
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
  },
};
