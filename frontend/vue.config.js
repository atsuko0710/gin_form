module.exports = {
    assetsDir: "static",
    devServer: {
        proxy: {
            '/api/v1': {
              target: 'http://127.0.01:8080',
              changeOrigin: true,
            }
        }
    }
  }
