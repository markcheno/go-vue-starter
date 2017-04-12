<template>
<div id="app">
  <nav class="navbar navbar-default">
    <div class="container">
      <ul class="nav navbar-nav">
        <li><router-link to="home">Home</router-link></li>
        <li v-if="!isAuthenticated()"><router-link to="login">Login</router-link></li>
        <li v-if="!isAuthenticated()"><router-link to="signup">Sign Up</router-link></li>
        <li v-if="isAuthenticated()"><router-link to="secretquote">Secret Quote</router-link></li>
        <li v-if="isAuthenticated()"><router-link to="logout" @click.native="logout()">Logout</router-link></li>
      </ul>
    </div>
  </nav>
  <div class="container">
    <router-view></router-view>
  </div>
</div>
</template>

<script>
import auth from './auth'

export default {

  data () {
    return {
    }
  },

  methods: {
    logout () {
      auth.logout(this)
    },
    isAuthenticated () {
      return auth.isAuthenticated()
    }
  }
}
</script>
