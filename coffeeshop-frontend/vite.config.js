import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), vueDevTools()],
  server: {
    hmr: true,
    proxy: {
      // 将 `/api` 路径的请求代理到后端服务器
      "/api": {
        target: "http://localhost:8080", // 后端服务器地址
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, '') // 移除 /api 前缀
      },
    },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
