const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: './src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
    clean: true,
  },
  module: {
    rules: [
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
      {
        test: /\.html$/i,
        loader: 'html-loader',
      },
      {
        test: /\.(png|jpe?g|gif|svg)$/i,
        type: 'asset/resource',
        generator: {
          filename: 'src/img/[name][ext]'
        }
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './src/pages/splash.html',
      filename: 'splash.html',
    }),
    new HtmlWebpackPlugin({
      template: './src/pages/pin.html',
      filename: 'pin.html',
    }),
    new HtmlWebpackPlugin({
      template: './src/pages/bank_main.html',
      filename: 'bank_main.html',
    }),
  ],
  devServer: {
    static: './dist',
    open: true,
    port: 3000,
  },
  mode: 'development',
};