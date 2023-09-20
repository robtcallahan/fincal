import { Layout, Register } from '@/views/register';

export default {
    path: '/register',
    component: Layout,
    children: [
        { path: '', component:  Register}
    ]
};