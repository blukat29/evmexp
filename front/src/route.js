import { createRouter, createWebHashHistory } from 'vue-router'

import Search from './pages/Search.vue'
import Code from './pages/Code.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', component: Search },
    { path: '/code/bin/:id', component: Code,
      props: { onchain: false } },
    { path: '/code/addr/:id', component: Code,
      props: { onchain: true } },
  ]
});

export default router;
