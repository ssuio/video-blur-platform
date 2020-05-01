module.exports = {
    presets: [
      "@babel/env",
      "@babel/preset-react"
    ],
    sourceType: "unambiguous",
    plugins: [
      "@babel/plugin-transform-runtime",
      ["@babel/plugin-proposal-decorators", { legacy: true }],
      ["@babel/plugin-proposal-class-properties", { loose: true }],
      "react-hot-loader/babel",
      [
        "module-resolver",
        {
          extensions: ['.js', '.jsx'],
          root: ['./src'],
          alias: {
            "components/*": ["src/components/*"],
            "containers/*": ["src/containers/*"],
            "assets/*": ["src/assets/*"],
            "static/*": ["src/static/*"],
            "theme/*": ["src/theme/*"]
          }
        }
      ]
    ]
  };
  