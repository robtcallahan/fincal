import {createApp} from '@vue/compat'
import { createPinia } from 'pinia';

import App from './App.vue'
import { router } from './router';

import {BootstrapVue, BTable} from 'bootstrap-vue'
import {BootstrapIconsPlugin} from "bootstrap-icons-vue";

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

const app = createApp(App);
app
    .use(BootstrapVue)
    .use(BootstrapIconsPlugin)
    .use(createPinia())
    .use(router)
    .component('b-table', BTable)
    .mount('#app');

