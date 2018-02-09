import Vuex from 'vuex'
import Vue from 'vue'
import { HTTP } from '../api'

const state = {
  isLoggedIn: false,
  user: {},
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
  ]
}

const actions = {
  logout({ commit }) {
    localStorage.removeItem('token')
    commit('setIsLoggedIn', false)
    commit('setUser', {})
  },
  setIsLoggedIn({ commit }, payload) {
    commit('setIsLoggedIn', payload)
  },
  loadProfile({ commit }, username) {
    HTTP.get('user/' + username)
      .then(response => {
        commit('setUser', response.data)
      })
      .catch(e => {
        console.log(e)
      })
  }
}

const mutations = {
  setIsLoggedIn(state, payload) {
    state.isLoggedIn = payload
  },
  setUser(state, payload) {
    state.user = payload
  }
}

const getters = {
  isLoggedIn(state) {
    return state.isLoggedIn
  },
  user(state) {
    return state.user
  },
  jobCategories(state) {
    return state.jobCategories
  },
  coinLabels(state) {
    return state.coinLabels
  }
}

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  actions,
  mutations,
  getters
})
