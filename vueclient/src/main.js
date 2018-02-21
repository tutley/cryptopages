import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VList,
  VBtn,
  VIcon,
  VChip,
  VCard,
  VCheckbox,
  VToolbar,
  VFooter,
  VDivider,
  VForm,
  VJumbotron,
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
    VChip,
    VIcon,
    VCard,
    VCheckbox,
    VToolbar,
    VJumbotron,
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
    primary: '#8BC34A',
    secondary: '#689F38',
    accent: '#FFEB3B',
    error: '#FF5252',
    info: '#73ea7b',
    success: '#DCEDC8',
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
