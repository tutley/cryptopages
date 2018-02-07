<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm6 offset-sm3>
        <v-card>
          <v-progress-linear :indeterminate="true" v-show="loading"></v-progress-linear>
          <v-card-title headline>
            <h2>{{ user.username }}</h2>
          </v-card-title>
          <v-card-text v-if="!editable">
            <!-- show all the user info that is available -->
            Lots of user details.
          </v-card-text>
          <v-card-text v-if="editable">
            <!-- show the user info as a form that can be edited -->
            <v-form v-model="valid" ref="form" lazy-validation>
              <v-text-field
                label="Name"
                v-model="user.Name"
                :rules="nameRules"
              ></v-text-field>
              <v-checkbox label="Available? (included in search results)" v-model="user.available"></v-checkbox>
              <v-text-field
                label="Email"
                v-model="user.email.value"
                :rules="emailRules"
                validate-on-blur
                hint="Let others know the best email to reach you with. You have the option to hide this."
              ></v-text-field>
              <v-checkbox label="Display Email?" v-model="user.email.makePublic"></v-checkbox>
            </v-form>
            </v-card-text>
            <v-card-actions v-show="!editable && thisUser">
              <v-btn small color="accent" @click="doEdit(user)">EDIT</v-btn>
              <v-spacer></v-spacer>
              <v-btn small color="error" @click="doDelete">DELETE</v-btn>
            </v-card-actions>
            <v-card-actions v-show="editable">
              <v-btn small color="success" @click="saveUpdate">SAVE</v-btn>
              <v-spacer></v-spacer>
              <v-btn small color="warning" @click="cancelEdit">CANCEL</v-btn>
            </v-card-actions>
            <v-progress-linear :indeterminate="true" v-show="processing"></v-progress-linear>
        </v-card>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
    <v-dialog v-model="showDelete">
      <v-card>
        <v-card-title class="headline">Really delete user?</v-card-title>
        <v-card-text>Are you sure? Think of the children!</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" flat @click="cancelDelete">CANCEL</v-btn>
          <v-btn color="error" flat @click="sendDelete">DELETE</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  data: () => ({
    user: {},
    errors: [],
    editable: false,
    showDelete: false,
    _originaluser: {},
    loading: false,
    processing: false
  }),
  computed: {
    thisUser() {
      return this.user.username === this.loggedInUser.username
    },
    loggedInUser() {
      return this.$store.getters.user
    }
  },
  methods: {
    doEdit: function(user) {
      this._originaluser = Object.assign({}, user)
      this.editable = true
    },
    cancelEdit: function() {
      Object.assign(this.user, this._originaluser)
      this.editable = false
    },
    saveUpdate: function() {
      if (this.$refs.form.validate()) {
        this.processing = true
        HTTP.put('user/' + this.$route.params.username, this.user)
          .then(response => {
            this.editable = false
            this.processing = false
          })
          .catch(e => {
            this.errors.push(e)
            this.processing = false
          })
      }
    },
    doDelete: function() {
      this.showDelete = true
    },
    cancelDelete: function() {
      this.showDelete = false
    },
    sendDelete: function() {
      this.processing = true
      HTTP.delete('user/' + this.$route.params.username)
        .then(response => {
          this.processing = false
          this.$router.push({ name: 'Home' })
        })
        .catch(e => {
          this.errors.push(e)
          this.processing = false
        })
    }
  },
  created() {
    this.loading = true
    HTTP.get('user/' + this.$route.params.username)
      .then(response => {
        this.user = response.data
        this.loading = false
      })
      .catch(e => {
        this.errors.push(e)
        this.loading = false
      })
  }
}
</script>
