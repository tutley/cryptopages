import Vuex from 'vuex'
import Vue from 'vue'

const state = {
  isLoggedIn: false
}

const actions = {
  logout({ commit }) {
    localStorage.removeItem('token')
    commit('setIsLoggedIn', false)
  },
  setIsLoggedIn({ commit }, payload) {
    commit('setIsLoggedIn', payload)
  }
}

const mutations = {
  setIsLoggedIn(state, payload) {
    state.isLoggedIn = payload
  }
}

const getters = {
  isLoggedIn(state) {
    return state.isLoggedIn
  }
}

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  actions,
  mutations,
  getters
})
