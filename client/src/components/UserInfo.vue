<template>
  <div class="col-sm-6 col-sm-offset-3">
    <h1>User Info</h1>
    <pre>{{ userinfo }}</pre>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  data () {
    return {
      userinfo: ''
    }
  },
  mounted () {
    this.getUserInfo()
  },
  methods: {
    getUserInfo () {
      this.$http.get('/api/user/info', { headers: auth.getAuthHeader() }).then(
        response => {
          this.userinfo = response.body
        },
        response => {
          if (response.status === 401) {
            auth.logout(this)
          }
          console.log(response)
        }
      )
    }
  }
}
</script>
