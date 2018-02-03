<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm8 offset-sm2>
        <p class="display-2">Sign Up!</p>
        <p class="body-1">Provide as much or as little information as you wish. You must create 
          an account to be able to search and view listings, but you don't have to make yourself 
          visible to others. Simply check "available" if you want others to be able to find you!</p>
        <v-form v-model="valid" ref="form" lazy-validation>
          <v-text-field
            label="Username"
            v-model="user.username"
            :rules="usernameRules"
            :counter="60"
            hint="This is your identifier on this site. It must be unique, and others will see it."
            @blur="checkUsername"
            required
          ></v-text-field>
          <v-text-field
            label="Password"
            v-model="user.password"
            :rules="passwordRules"
            required
            type="password"
          ></v-text-field>
          <v-checkbox label="Available? (included in search results)" v-model="user.available"></v-checkbox>
          <v-text-field
            label="Name"
            v-model="user.name.value"
            :rules="nameRules"
            hint="Put your name here."
          ></v-text-field>
          <v-text-field
            label="Email"
            v-model="user.email.value"
            :rules="emailRules"
            hint="Let others know the best email to reach you with. You have the option to hide this."
          ></v-text-field>
          <v-checkbox label="Display Email?" v-model="user.email.makePublic"></v-checkbox>
           <v-select
              :items="jobCategories"
              v-model="user.jobCategory"
              label="Job Category"
              single-line
              item-text="name"
              item-value="val"
              return-object
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
              <v-flex class="xs4 sm3" v-for="(item, key) in user.coins"       :key="key">
                <v-checkbox 
                  :label="coinLabels[key]"
                  v-model="user.coins[key]"></v-checkbox>
              </v-flex>
            </v-layout>

          <v-btn
            @click="submit"
            :disabled="!valid"
          >
            Submit
          </v-btn>
          <v-btn @click="clear">Clear</v-btn>
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
    usernameAvailable: true,
    user: {
      username: '',
      password: '',
      available: true,
      name: '',
      email: {
        value: '',
        makePublic: true
      },
      location: {
        value: '',
        makePublic: true
      },
      skills: '',
      jobCategory: '',
      jobDescription: '',
      coins: {
        bcc: false,
        btc: false,
        eth: false,
        ltc: false,
        neo: false,
        other: false,
        xlm: false,
        xrp: false
      },
      otherCoin: ''
    },
    coinLabels: {
      bcc: 'Bitcoin Cash',
      btc: 'Bitcoin',
      eth: 'Ethereum',
      ltc: 'Litecoin',
      neo: 'Neo',
      other: 'Other',
      xlm: 'Lumen',
      xrp: 'Ripple'
    },
    jobCategories: [
      { name: 'Hardware', val: 'hardware' },
      { name: 'Software', val: 'software' },
      { name: 'Writing', val: 'writing' },
      { name: 'Legal', val: 'legal' },
      { name: 'Labor', val: 'labor' },
      { name: 'Automotive', val: 'automotive' },
      { name: 'Services', val: 'services' },
      { name: 'Others', val: 'others' }
    ],
    usernameRules: [
      v => !!v || 'Username is required',
      v => (v && v.length <= 60) || 'Username must be less than 60 characters',
      v => this.usernameAvailable || 'Username is not available, sorry'
    ],
    passwordRules: [v => !!v || 'Password is required'],
    nameRules: [],
    emailRules: [],
    skillsRules: [],
    jobDescriptionRules: [],
    errors: [],
    valid: true
  }),
  methods: {
    checkUsername(un) {
      // check the value of this.user.username to see if the user already exists in the db
      HTTP.get('user/checkUsername/' + this.user.username)
        .then(response => {
          this.usernameAvailable = false
        })
        .catch(e => {
          this.usernameAvailable = true
        })
    },
    submit() {
      let user = this.user
      user.skills = this.user.skills.split(' ')
      if (this.$refs.form.validate()) {
        this.sending = true
        HTTP.post('/user', user)
          .then(response => {
            this.sending = false
            console.log(response.data)
            this.$router.push({ name: 'User Detail', params: { id: response.data.username } })
          })
          .catch(e => {
            this.sending = false
            this.errors.push(e)
          })
      }
    },
    clear() {
      this.$refs.form.reset()
    }
  }
}
</script>
