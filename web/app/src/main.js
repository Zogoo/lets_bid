import "@babel/polyfill";
import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import router from "./routes";
import vuetify from './plugins/vuetify';

Vue.config.productionTip = process.env.NODE_ENV == "production";

// router.beforeEach(Vue.prototype.$auth.authRedirectGuard());

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App)
}).$mount("#app");
