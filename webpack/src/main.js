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

Vue.config.productionTip = false

import auth from './auth'

function requireAuth (route, redirect, next) {
  if (!auth.user.authenticated) {
    console.log('requireAuth: not authenticated')
    redirect({
      path: '/login',
      query: { redirect: route.fullPath }
    })
  } else {
    console.log('requireAuth: authenticated')
    next()
  }
}

const router = new VueRouter({
  // mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/logout',
      beforeEnter (route, redirect) {
        auth.logout()
        redirect('/')
      }
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
    {
      path: '/secretquote',
      name: 'secretquote',
      component: SecretQuote,
      beforeEnter: requireAuth
    },
    {
      path: '*',
      redirect: {
        name: 'home'
      }
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
