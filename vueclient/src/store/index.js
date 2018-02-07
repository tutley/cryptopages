import Vuex from 'vuex'
import Vue from 'vue'
import { HTTP } from '../api'

const state = {
  isLoggedIn: false,
  user: {}
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
  }
}

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  actions,
  mutations,
  getters
})
