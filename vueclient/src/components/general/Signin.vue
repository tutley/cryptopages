<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm8 offset-sm2>
        <p class="headline">Sign In</p>
        <v-form  v-model="valid" ref="form" lazy-validation>
          <v-text-field
            label="Username"
            autofocus
            v-model="username"
            :rules="usernameRules"
            required
          ></v-text-field>
          <v-text-field
            label="Password"
            v-model="password"
            :rules="passwordRules"
            @keyup.enter="submit"
            required
            type="password"
          ></v-text-field>
          <v-btn
            @click="submit"
            :disabled="!valid"
          >
            Submit
          </v-btn>
        </v-form>
        <v-progress-linear :indeterminate="true" v-show="sending"></v-progress-linear>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  data: () => ({
    sending: false,
    username: '',
    usernameRules: [v => !!v || 'Username is required'],
    password: '',
    passwordRules: [v => !!v || 'Password is required'],
    errors: [],
    valid: true
  }),
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        this.sending = true
        HTTP.post(
          '/jwt/signin',
          {},
          {
            auth: {
              username: this.username,
              password: this.password
            }
          }
        )
          .then(response => {
            this.sending = false
            let token = response.headers.authorization.toString().replace('Bearer ', '')
            localStorage.setItem('token', token)
            HTTP.defaults.headers['Authorization'] = 'Bearer ' + token
            this.$store.dispatch('setIsLoggedIn', true)
            this.$store.dispatch('loadProfile', this.username)
            this.$router.push({ name: 'Hello' })
          })
          .catch(e => {
            this.sending = false
            this.errors.push(e)
          })
      }
    }
  }
}
</script>

<style>

</style>
