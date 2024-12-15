import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {copy} from 'vite-plugin-copy'
// https://vite.dev/config/
export default defineConfig({
    plugins: [
        vue(),
    ],
    publicDir: 'public',
    build: {
        assetsDir: 'assets',
    },
    server: {
        host: '0.0.0.0',
        port: 3090
    },
    preview: {
        host: '0.0.0.0',
        port: 3090 // 指定预览模式的端口
    },

})
