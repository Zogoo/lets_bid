import "@babel/polyfill";
import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import router from "./routes";
import vuetify from './plugins/vuetify';
import VueFlashMessage from "vue-flash-message";
import Vuex from "vuex";

Vue.use(Vuex);
Vue.use(VueFlashMessage);

require("vue-flash-message/dist/vue-flash-message.min.css");

Vue.config.productionTip = process.env.NODE_ENV == "production";

// router.beforeEach(Vue.prototype.$auth.authRedirectGuard());
// TODO: move this part to different package
Vue.prototype.$auth = {
  accessTokenField: "access-token",
  setAccessToken(token) {
    localStorage.setItem(this.accessTokenField, token);
  },
  clearAccessToken() {
    localStorage.setItem(this.accessTokenField, null);
  },
  getAccessToken() {
    let storedToken = localStorage.getItem(this.accessTokenField);
    if (storedToken != null) {
      return this.parseJwt(storedToken);
    } else {
      return null
    }
  },
  checkAuthenticated() {
    try {
      this.parseJwt(localStorage.getItem(this.accessTokenField));
      return true;
    } catch (error) {
      return false;
    }
  },
  parseJwt(token) {
    let base64Url = token.split(".")[1];
    let base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    let jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map(function(c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );

    return JSON.parse(jsonPayload);
  }
};

router.beforeEach((to, from, next) => {
  // redirect to signin page if not logged in and trying to access a restricted page
  const authRequired = to.matched.some(record => record.meta.authRequired);
  const loggedIn = Vue.prototype.$auth.checkAuthenticated();

  if (authRequired && !loggedIn) {
    return next("/signin");
  }

  next();
});

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App)
}).$mount("#app");
