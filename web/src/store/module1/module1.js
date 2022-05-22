import { getField, updateField } from 'vuex-map-fields';

export const state = () => ({
    posts: []
})

const getters = {
    getField,
    posts: (state) => state.posts,
    postByID(state) {
        return id => state.posts.filter(post => post.id === id)
    },
}

const mutations = {
    updateField,
    SET_POSTS: (state, posts) => {
        state.posts = posts
    },
    CLEAR_POSTS: (state) => {
        state.posts = []
    }
}

const actions = {
    setPosts({ commit, state }, obj) {
        commit('SET_POSTS', obj)
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
}
