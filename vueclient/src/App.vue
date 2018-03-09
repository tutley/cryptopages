<template>
<v-app id="cryptopages">
    <v-navigation-drawer
      app
      fixed
      v-model="drawer"
      light
    >
      <v-list>
        <v-list-tile
          v-for="item in menuItems"
          :key="item.title"
          exact
          :to="item.link">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>{{ item.title }}</v-list-tile-content>
        </v-list-tile>
        <v-list-tile
          v-if="userIsAuthenticated"
          @click="onLogout">
          <v-list-tile-action>
            <v-icon>exit_to_app</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>Logout</v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app dark color="primary">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>
        Cryptopages
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn flat to="/about">About</v-btn>
     </v-toolbar>
      <v-content ref="content">
          <v-btn fab color="accent" 
            bottom right fixed small
            @click="jump('/search')">
            <v-icon>search</v-icon>
          </v-btn>
        <div>&nbsp;</div>
        <router-view></router-view>
         <v-footer class="pa-3">
          <div><a href="https://tomutley.com/">Tom Utley</a></div>
          <v-spacer></v-spacer>
        </v-footer>
      </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'app',
  data() {
    return {
      drawer: false
    }
  },
  computed: {
    userIsAuthenticated() {
      return this.$store.getters.isLoggedIn
    },
    user() {
      return this.$store.getters.user
    },
    menuItems() {
      let menuItems = [
        { icon: 'home', title: 'Home', link: '/' },
        { icon: 'info', title: 'About', link: '/about' },
        { icon: 'lock_open', title: 'Sign in', link: '/signin' },
        { icon: 'add_circle', title: 'Sign Up', link: '/signup' }
      ]
      if (this.userIsAuthenticated) {
        menuItems = [
          { icon: 'home', title: 'Home', link: '/' },
          { icon: 'info', title: 'About', link: '/about' },
          { icon: 'search', title: 'Search Listings', link: '/search' }
        ]
      }
      if (this.user.username) {
        menuItems.push({
          icon: 'account_circle',
          title: 'My Profile',
          link: '/user/' + this.user.username
        })
      }
      return menuItems
    }
  },
  methods: {
    jump(loc) {
      this.$router.push(loc)
    },
    onLogout() {
      this.$store.dispatch('logout')
      this.$router.push('/')
    }
  },
  beforeMount() {
    // This removes the spinner once the app is loaded
    let pwadiv = document.getElementById('pwaloader')
    if (pwadiv) {
      pwadiv.remove()
    }
  }
}
</script>

<style lang="stylus">
$color-pack = false;

@import '~vuetify/src/stylus/main';
</style>
