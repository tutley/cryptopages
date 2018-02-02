import '../node_modules/vuetify/src/stylus/app.styl'
import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VList,
  VBtn,
  VIcon,
  VCard,
  VCheckbox,
  VToolbar,
  VFooter,
  VDivider,
  VForm,
  VProgressCircular,
  VProgressLinear,
  VSelect,
  VSubheader,
  VTextField,
  VAlert,
  VGrid,
  VDialog,
  transitions
} from 'vuetify'
import App from './App'
import router from './router'
import store from './store/index'

Vue.config.productionTip = false

Vue.use(Vuetify, {
  components: {
    VApp,
    VNavigationDrawer,
    VList,
    VBtn,
    VIcon,
    VCard,
    VCheckbox,
    VToolbar,
    VFooter,
    VDivider,
    VForm,
    VProgressCircular,
    VProgressLinear,
    VSelect,
    VSubheader,
    VTextField,
    VAlert,
    VGrid,
    VDialog,
    transitions
  },
  directives: {
    Touch
  },
  theme: {
    primary: '#ffdb3b',
    secondary: '#ffab0b',
    accent: '#757575',
    error: '#FF5252',
    info: '#73ea7b',
    success: '#4CAF50',
    warning: '#FFC107'
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App },
  created() {
    if (localStorage.getItem('token')) {
      this.$store.dispatch('setIsLoggedIn', true)
    }
  }
})
