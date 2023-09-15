import {createApp} from '@vue/compat'
import {BootstrapVue, BTable} from 'bootstrap-vue'
import {BootstrapIconsPlugin} from "bootstrap-icons-vue";
import App from './App.vue'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// Vuetify
// import 'vuetify/styles'
// import {createVuetify} from 'vuetify'
// import * as components from 'vuetify/components'
// import * as directives from 'vuetify/directives'
//
// const vuetify = createVuetify({
//     components,
//     directives,
// })

const app = createApp(App);
app
    .use(BootstrapVue)
    .use(BootstrapIconsPlugin)
    .component('b-table', BTable)
    .mount('#app');

