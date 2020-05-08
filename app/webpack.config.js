const path = require("path");
const fs = require("fs");
const HtmlWebPackPlugin = require("html-webpack-plugin");

module.exports = {
  output: {
    publicPath: "/",
    path: path.resolve(__dirname, '../server/dist')
  },
  module: {
    rules: [
      {
        test: /\.js(x)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader",
            options: {
              plugins: ["react-hot-loader/babel"],
            },
          },
        ],
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.scss$/,
        use: ["style-loader", "css-loader", "postcss-loader", "sass-loader"],
      },
      {
        test: /\.(png|jpg|gif|mp4|ogg|svg|woff|woff2|ttf|eot)$/,
        loader: "file-loader",
      },
    ],
  },
  resolve: {
    alias: {
      components: path.resolve(__dirname, "src/components/"),
      containers: path.resolve(__dirname, "src/containers/"),
      assets: path.resolve(__dirname, "src/assets/"),
      static: path.resolve(__dirname, "src/static/"),
      theme: path.resolve(__dirname, "src/theme/"),
      "react-dom": "@hot-loader/react-dom",
    },
    extensions: [".js", ".jsx"],
    modules: [path.resolve(__dirname, "src/"), "node_modules/"],
  },
  optimization: {
    splitChunks: {
      chunks: "all",
    },
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: path.resolve(__dirname, "src/index.html"),
      filename: "./index.html",
    }),
  ],
  devServer: {
    compress: true,
    historyApiFallback: true,
    hot: true,
    disableHostCheck: true
  },
};
