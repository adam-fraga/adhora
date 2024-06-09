const path = require("path"); // Node Js module path management

module.exports = {
  devtool: "eval-source-map",
  entry: "./src/script/index.ts",
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: "ts-loader",
        include: [path.resolve(__dirname, "src/script")],
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: [".js", ".ts"],
  },
  output: {
    publicPath: "src/script", // Specify the folder for webpack dev server to reload properly
    filename: "bundle.js", // Name of the outputed file after compile
    path: path.resolve(__dirname, "static/js"), // Abs path for the outputed file
  },
};
