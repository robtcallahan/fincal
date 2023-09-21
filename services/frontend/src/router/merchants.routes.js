import { Layout, Merchants } from '@/views/merchants';

export default {
    path: '/merchants',
    component: Layout,
    children: [
        { path: '', component:  Merchants}
    ]
};