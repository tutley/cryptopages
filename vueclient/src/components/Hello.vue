<template>
<v-container fluid>
  <v-layout row>
    <v-flex xs10 offset-xs1>
      <v-jumbotron color="grey lighten-2">
        <v-container fill-height>
          <v-layout align-center>
            <v-flex>
              <h3 class="display-3">Welcome to Cryptopages</h3>
              <span class="subheading"></span>
              <v-divider class="my-3"></v-divider>
              <div class="title mb-3">Sign up or sign in to get started!</div>
              <v-btn v-if="isLoggedIn" large color="primary" class="mx-0" to="/search">Search</v-btn>
              <v-btn v-if="!isLoggedIn" large color="primary" class="mx-0" to="/signin">Sign In</v-btn>
              <v-btn v-if="!isLoggedIn" large color="secondary" class="mx-0" to="/signup">Sign Up</v-btn>
            </v-flex>
          </v-layout>
        </v-container>
      </v-jumbotron>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
export default {
  name: 'hello',
  data() {
    return {
      msg: 'Its like the Yellow Pages but for Crypto!'
    }
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    }
  }
}
</script>

<style>
.container {
  padding: 0px;
}
.hello {
  padding: 12px;
}
</style>
