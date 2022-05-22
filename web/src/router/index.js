import Vue from 'vue'
import VueRouter from 'vue-router'
import BookListing from '../views/BookListing.vue'
import BorrowedListing from '../views/BorrowedListing.vue'
import About from '../views/About.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/books',
  },
  {
    path: '/books',
    name: 'BookListing',
    component: BookListing,
  },
  {
    path: '/borrowed-list',
    name: 'BorrowedListing',
    component: BorrowedListing,
  },
  {
    path: '/about',
    name: 'About',
    component: About,
  }
]

const router = new VueRouter({
  routes
})

export default router
