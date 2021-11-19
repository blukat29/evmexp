import { createApp } from 'vue'
import { Quasar } from 'quasar'
import quasarUserOptions from './quasar-user-options'

import App from './App.vue'
import router from './route.js'
import installMock from './mock/mock.js'

installMock();

const app = createApp(App);
app.use(router);
app.use(Quasar, quasarUserOptions);
app.mount('#app');
