<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm6 offset-sm3>
        <h4>Need to refactor this page to be a search page and to show a list of search results</h4>
        <!-- <v-progress-linear :indeterminate="true" v-show="loading"></v-progress-linear>
        <v-alert v-show="users.length < 1 && !loading" color="info" icon="info" value="true">
        No users yet...   (click that little green add button)
        </v-alert>
        <v-card v-show="errors.length < 1">
          <v-list three-line>
            <v-subheader>users</v-subheader>
            <template v-for="(user, index) in users">
              <v-divider :key="index"></v-divider>
              <v-list-tile v-bind:key="index" @click="displayDetails(user.id)">
                <v-list-tile-content>
                  <v-list-tile-title v-html="user.title"></v-list-tile-title>
                  <v-list-tile-sub-title v-html="user.body"></v-list-tile-sub-title>
                </v-list-tile-content>
              </v-list-tile>
            </template>
          </v-list>
        </v-card>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert> -->
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  methods: {
    displayDetails(username) {
      this.$router.push({ name: 'User Detail', params: { username: username } })
    }
  },
  data: () => ({
    userss: [],
    errors: [],
    loading: false
  }),
  created() {
    this.loading = true
    HTTP.get('user')
      .then(response => {
        this.users = response.data
        this.loading = false
      })
      .catch(e => {
        this.errors.push(e)
        this.loading = false
      })
  }
}
</script>
