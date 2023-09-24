import {fileURLToPath, URL} from 'url';

import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            vue: '@vue/compat',
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    optimizeDeps: {exclude: ["fsevents"]},
    plugins: [
        vue({
            template: {
                compilerOptions: {
                    compatConfig: {
                        MODE: 3
                    }
                }
            }
        })
    ],
    server: {
        host: true,
        port: 80,
        watch: {
            usePolling: true
        }
    }
})