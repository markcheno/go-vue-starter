// import router from '../router'

const API_URL = 'http://localhost:3000/'
const LOGIN_URL = API_URL + 'login'
const SIGNUP_URL = API_URL + 'signup'

export default {

  user: {
    authenticated: false
  },

  login (context, creds, redirect) {
    console.log('login: creds=', creds)
    context.$http.post(LOGIN_URL, creds).then(response => {
      console.log('login: response=', response)
      console.log('login: token=', response.body.id_token)
      localStorage.setItem('id_token', response.body.id_token)
      this.user.authenticated = true
      // if (redirect) {
      //   router.go(redirect)
      // }
    }, response => {
      context.error = response.statusText
    })
  },
  signup (context, creds, redirect) {
    console.log('auth: creds=', creds)
    context.$http.post(SIGNUP_URL, creds).then(response => {
      localStorage.setItem('id_token', response.body.id_token)
      this.user.authenticated = true
      // if (redirect) {
      //   router.go(redirect)
      // }
    }, response => {
      context.error = response.statusText
    })
  },
  logout () {
    console.log('logging out')
    localStorage.removeItem('id_token')
    this.user.authenticated = false
  },
  checkAuth () {
    var jwt = localStorage.getItem('id_token')
    if (jwt) {
      this.user.authenticated = true
    } else {
      this.user.authenticated = false
    }
  },
  getAuthHeader () {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('id_token')
    }
  }
}
