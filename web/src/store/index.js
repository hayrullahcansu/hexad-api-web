// import Vue from 'vue'
// import Vuex from 'vuex'

// Vue.use(Vuex)

// export default new Vuex.Store({
//   state: {
//   },
//   mutations: {
//   },
//   actions: {
//   },
//   modules: {
//   }
// })
import Vue from 'vue'
import Vuex from "vuex";
import module1 from "./module1/module1"
import libraryStore from "./library/libraryStore";

Vue.use(Vuex)

const createStore = () => {
  return new Vuex.Store({
    namespaced: true,
    modules: {
      module1: module1,
      libraryStore: libraryStore,
    }
  });
};

export default createStore;