import { defineConfig } from 'vite'  // 追加
import { fileURLToPath, URL } from 'node:url'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

export default defineConfig({  // 変更
  plugins: [
    vue(),
    vueDevTools(),
  ],
  server: {
    host: "0.0.0.0",
    port: 3000
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
