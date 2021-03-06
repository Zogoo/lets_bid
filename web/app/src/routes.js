import Vue from 'vue';
import VueRouter from 'vue-router';

import SignIn from "./components/SignIn";
import SignUp from "./components/SignUp"
import MyPage from './components/MyPage'

Vue.use(VueRouter)

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: SignIn },
    { path: "/cas", component: SignIn },
    { path: "/cas/signin", component: SignIn },
    { path: "/cas/signup", component: SignUp },
    { path: "/sso/my_page", component: MyPage, meta: { authRequired: true } }
  ]
});