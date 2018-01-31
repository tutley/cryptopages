<template>
  <v-container fluid>
    <v-layout>
      <v-flex>
          <div class="hello">
            <p class="subheading">
              Hello, welcome to cryptopages.
              </p>
              <p>
                This site is not ready to use yet. When it's ready, we'll update this page.
              </p>
             <!-- <center>
              <v-btn color="primary" to="/signup">Get Started</v-btn>
              </center> -->
          </div>
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
