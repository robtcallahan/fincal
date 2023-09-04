import vue from '@vitejs/plugin-vue'
// import vue from '@vue/compat'

// https://vitejs.dev/config/
export default {
  resolve: {
    alias: {
      vue: '@vue/compat'
    }
  },
  optimizeDeps: { exclude: ["fsevents"] },
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
  ]
}