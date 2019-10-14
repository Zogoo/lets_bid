import Vue from 'vue';
import Vuex from 'vuex';
import Client from './client';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    items: [],
  },
  mutations: {
    loadItems(state, items) {
      state.items = items
    },
    updateItem(state, item){
      state.items[item.id] = item;
    }
  },
  getters: {
    getItems(state){
      return state.items;
    }
  },
  actions: {
    initializeItems({commit}) {
      Client.getAllItems().then((data) => {
        commit('loadItems', data)
      })
    },
    updateItem({commit, state}, item) {
      if (state.items[item.id]) {
        return Client.updateItem(item.id, item).then((data) => {
          commit('updateItem', data)
        });
      }
    }
  }
});

export default store;