import { getField, updateField } from 'vuex-map-fields';
import { fetchWrapper } from '@/_helper'
import CONSTS from '@/constants/consts'


const state = () => ({
    booksLoading: false,
    books: [],
    borrowedList: [],
})

const getters = {
    getField,
    books: (state) => state.books,
    booksLoading: (state) => state.booksLoading,
    borrowedList: (state) => state.borrowedList,
}

const mutations = {
    updateField,
    SET_BOOKS: (state, books) => {
        state.books = books
    },
    SET_BOOKS_LOADING: (state, loading) => {
        state.booksLoading = loading
    },
    SET_BORROWED_LIST: (state, borrowedList) => {
        state.borrowedList = borrowedList
    },
}

const actions = {
    tryGetBooks({ commit }) {
        commit('SET_BOOKS_LOADING', true)
        fetchWrapper.get('v1/books').then(data => {
            console.log('response!', data)
            commit('SET_BOOKS_LOADING', false)
            if (data.Status === 200 && data.Body) {
                commit('SET_BOOKS', data.Body)
            }
        });
    },
    tryBorrowBook({ commit, dispatch }, reqBody) {
        commit('SET_BOOKS_LOADING', true)
        fetchWrapper
            .post('v1/books/borrow', reqBody)
            .then(data => {
                if (data.Status === 200 && data.Body) {
                    dispatch('tryGetBooks');
                } else {
                    alert(data.Body)
                    commit('SET_BOOKS_LOADING', false)
                }
            });
    },
    tryGetMyBorrowedList({ commit }, userName) {
        commit('SET_BOOKS_LOADING', true)
        fetchWrapper.get('v1/borrowed/' + userName).then(data => {
            console.log('response!', data)
            commit('SET_BOOKS_LOADING', false)
            if (data.Status === 200 && data.Body) {
                commit('SET_BORROWED_LIST', data.Body)
            }
        });
    },
    tryReturnBook({ commit, dispatch }, reqBody) {
        commit('SET_BOOKS_LOADING', true)
        fetchWrapper
            .post('v1/books/return', reqBody)
            .then(data => {
                if (data.Status === 200 && data.Body) {
                    dispatch('tryGetMyBorrowedList', CONSTS.USER);
                } else {
                    alert(data.Body)
                    commit('SET_BOOKS_LOADING', false)
                }
            });
    },
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
}
