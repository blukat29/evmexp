import { createRouter, createWebHashHistory } from 'vue-router'

import Search from './pages/Search.vue'
import Contract from './pages/Contract.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', component: Search },
    { path: '/contract', component: Contract },
  ]
});

export default router;
