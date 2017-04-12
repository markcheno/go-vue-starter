// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import Home from '@/components/Home'
import Login from '@/components/Login'
import Signup from '@/components/Signup'
import SecretQuote from '@/components/SecretQuote'

import VueResource from 'vue-resource'
Vue.use(VueResource)

import VueRouter from 'vue-router'
Vue.use(VueRouter)

Vue.config.productionTip = true

import auth from './auth'

function requireAuth (route, redirect, next) {
  if (!auth.user.authenticated) {
    // console.log('requireAuth: not authenticated')
    this.$router.replace('/login')
  } else {
    // console.log('requireAuth: authenticated')
    next()
  }
}

const router = new VueRouter({
  // mode: 'history',
  // base: __dirname,
  routes: [
    {
      path: '/',
      component: Home,
      beforeEnter: auth.checkAuth()
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
      beforeEnter: auth.checkAuth()
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      beforeEnter: auth.checkAuth()
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup,
      beforeEnter: auth.checkAuth()
    },
    {
      path: '/secretquote',
      name: 'secretquote',
      component: SecretQuote,
      beforeEnter: requireAuth
    }
  ]
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
