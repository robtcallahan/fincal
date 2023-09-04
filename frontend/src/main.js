import {configureCompat, createApp} from '@vue/compat'
import {BootstrapVue, BTable} from 'bootstrap-vue'
import {BootstrapIconsPlugin} from "bootstrap-icons-vue";

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

configureCompat({
    // MODE: 3,
    WATCH_ARRAY: false,
    RENDER_FUNCTION: 'suppress-warning',
    INSTANCE_LISTENERS: 'suppress-warning',
    COMPONENT_FUNCTIONAL: 'suppress-warning',
    OPTIONS_BEFORE_DESTROY: 'suppress-warning',
    INSTANCE_SCOPED_SLOTS: 'suppress-warning',
    OPTIONS_DATA_MERGE: 'suppress-warning',
    COMPONENT_V_MODEL: 'suppress-warning',
    CUSTOM_DIR: 'suppress-warning',
    INSTANCE_EVENT_EMITTER: 'suppress-warning',
    ATTR_FALSE_VALUE: 'suppress-warning',
    INSTANCE_ATTRS_CLASS_STYLE: 'suppress-warning',
    GLOBAL_PROTOTYPE: 'suppress-warning',
    GLOBAL_EXTEND: 'suppress-warning',
    GLOBAL_MOUNT: 'suppress-warning',
    OPTIONS_DESTROYED: 'suppress-warning',
    INSTANCE_DESTROY: 'suppress-warning',
});

import App from './App.vue'
const app = createApp(App)
app.use(BootstrapVue)
app.use(BootstrapIconsPlugin)
app.component('b-table', BTable)
// app.component('BIconXSquareFill', BIconXSquareFill)
app.mount('#app')
