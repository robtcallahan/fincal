import { Layout, Budget } from '@/views/budget';

export default {
    path: '/budget',
    component: Layout,
    children: [
        { path: '', component:  Budget}
    ]
};