import Vue, { createApp } from '@vue/compat'

import { BootstrapVue } from 'bootstrap-vue'
Vue.use(BootstrapVue)
import { BTable } from 'bootstrap-vue'
Vue.component('b-table', BTable)

import 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import App from './App.vue'
createApp(App).mount('#app')
