<template>
<div id="app">

  <nav class="navbar navbar-default navbar-static-top">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">Brand</a>
      </div>
      <div id="navbar" class="navbar-collapse collapse">
        <ul class="nav navbar-nav">
          <li><router-link to="home">Home</router-link></li>
          <li v-if="!isAuthenticated()"><router-link to="login">Login</router-link></li>
          <li v-if="!isAuthenticated()"><router-link to="signup">Sign Up</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="secretquote">Secret Quote</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="userinfo">User Info</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="logout" @click.native="logout()">Logout</router-link></li>
        </ul>
      </div>
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
