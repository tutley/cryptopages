<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm8 offset-sm2>
        <v-card>
          <v-progress-linear :indeterminate="true" v-show="loading"></v-progress-linear>
          <v-card-title headline>
            <h2>{{ user.username }}</h2>
          </v-card-title>
          <v-list two-line v-if="!editable">
            <v-list-tile>
              <!-- <v-list-tile-action>
                <v-icon color="indigo">phone</v-icon>
              </v-list-tile-action> -->
              <v-list-tile-content v-if="user.name">
                <v-list-tile-sub-title>Name</v-list-tile-sub-title>
                <v-list-tile-title>{{user.name}}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile v-if="user.email">
              <v-list-tile-content v-if="user.email.makePublic">
                <v-list-tile-sub-title>Email</v-list-tile-sub-title>
                <v-list-tile-title>{{user.email.value}}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile v-if="user.location">
              <v-list-tile-content v-if="user.location.makePublic">
                <v-list-tile-sub-title>Location</v-list-tile-sub-title>
                <v-list-tile-title>{{user.location.value}}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content>
                <v-list-tile-sub-title>Job Category</v-list-tile-sub-title>
                <v-list-tile-title>{{ jobCategories.find(x => x.val === user.jobCategory).name }}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content v-if="user.jobDescription">
                <v-list-tile-sub-title>Job Description</v-list-tile-sub-title>
                <v-list-tile-title>{{user.jobDescription}}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content v-if="user.skills.length > 0">
                <v-list-tile-sub-title>Skills</v-list-tile-sub-title>
                <v-list-tile-title>
                  {{user.name}}

                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content>
                <v-list-tile-sub-title>Coins Accepted</v-list-tile-sub-title>
                <v-list-tile-title>
                  {{user.name}}
                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            
          </v-list>
          <v-card-text v-if="editable">
            <!-- show the user info as a form that can be edited -->
            <v-form v-model="valid" ref="form" lazy-validation>
              <v-text-field
                label="Name"
                v-model="user.name"
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
              <v-text-field
                label="location"
                v-model="user.location.value"
                :rules="locationRules"
                validate-on-blur
                hint="Let others know where you are located. You have the option to hide this."
              ></v-text-field>
              <v-checkbox label="Display Location?" v-model="user.location.makePublic"></v-checkbox>
                 <v-select
                  :items="jobCategories"
                  v-model="user.jobCategory"
                  label="Job Category"
                  single-line
                  item-text="name"
                  item-value="val"
                  hint="What type of work are you interested in?"
                ></v-select>
              <v-text-field
                label="Job Description"
                v-model="user.jobDescription"
                :rules="jobDescriptionRules"
                hint="What type of work are you looking for? Full-time, part-time, contract, big projects, small projects, etc"
                ></v-text-field>
              <v-text-field
                label="Skills"
                v-model="user.skills"
                :rules="skillsRules"
                hint="What skills do you bring to the table? Use spaces to separate skills."
              ></v-text-field>
              <span class="subheading">Coins Accepted</span><br/>
                <v-layout row wrap>
                  <v-flex class="xs4 sm3" v-for="(item, key) in user.coins" :key="key">
                    <v-checkbox 
                      :label="coinLabels[key]"
                      v-model="user.coins[key]"></v-checkbox>
                  </v-flex>
                </v-layout>
              <v-text-field
              v-show="user.coins.other"
                label="Other Coin Name"
                v-model="user.otherCoin"
                :rules="otherCoinRules"
                hint="What is the name of the other crypto currency you accept?"
              ></v-text-field>
            </v-form>
            </v-card-text>
            <v-card-actions v-show="!editable && thisUser">
              <v-btn small color="secondary" @click="doEdit(user)">EDIT</v-btn>
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
    nameRules: [],
    emailRules: [
      v =>
        v.length === 0 ||
        /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) ||
        'E-mail must be valid'
    ],
    locationRules: [],
    skillsRules: [],
    jobDescriptionRules: [],
    otherCoinRules: [],
    errors: [],
    valid: true,
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
    },
    coinLabels() {
      return this.$store.getters.coinLabels
    },
    jobCategories() {
      return this.$store.getters.jobCategories
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
  beforeCreate() {
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
