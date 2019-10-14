import Vue from 'vue';
import VueRouter from 'vue-router';

import Login from './components/Login'
import MyPage from './components/MyPage'

Vue.use(VueRouter)

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: Login },
    { path: "/my_page", component: MyPage }
  ]
});