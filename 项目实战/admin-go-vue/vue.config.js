/**
 * vue配置
 * 
 */
const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  lintOnSave: false,
  productionSourceMap: false,
  publicPath:"/",
  outputDir:'dist',
  assetsDir:'assets',
  devServer: {
    port: 2003,
    host: '0.0.0.0',
    https: false,
    open: false,
    proxy: {
      "/api": {
        target: "http://localhost:2002",
        changeOrigin: true,
      }
    },
    client:{
      overlay: false // 关闭全屏报错 Uncaught runtime errors:
    }
  }
})
