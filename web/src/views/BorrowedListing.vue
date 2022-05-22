<template>
  <v-container class="grey lighten-5" fluid>
    <!-- Stack the columns on mobile by making one full-width and the other half-width     class="pa-xs-0 mx-lg-auto" -->
    <h1 class="text-center">My Borrowed List</h1>
    <v-row>
      <v-col cols="12" md="12">
        <v-card>
          <v-list v-if="borrowedList.length > 0">
            <v-list-item two-line v-for="book in borrowedList" :key="book.Name">
              <v-list-item-content>
                <v-list-item-title
                  >{{ book.Name }}
                  <div class="pa-1">
                    <v-btn block @click="returnBook(book.Name)"> Return </v-btn>
                  </div></v-list-item-title
                >
                <v-list-item-subtitle
                  >User: {{ book.User }}</v-list-item-subtitle
                >
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <div v-else>
            <h4 class="text-center">You have not borrowed any book yet.</h4>
          </div>
        </v-card>
        <v-progress-linear
          indeterminate
          color="cyan"
          v-if="booksLoading"
        ></v-progress-linear>
        <div v-else></div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// import HelloWorld from "../components/HelloWorld";

import { mapActions, mapGetters, mapMutations, mapState } from "vuex";
// import _ from "lodash";
import CONSTS from "../constants/consts";

export default {
  name: "BorrowedListing",
  created() {
    this.tryGetMyBorrowedList(CONSTS.USER);
  },
  computed: {
    ...mapGetters("libraryStore", ["booksLoading", "borrowedList"]),
  },
  methods: {
    ...mapActions("libraryStore", ["tryGetMyBorrowedList", "tryReturnBook"]),
    returnBook(book) {
      const reqBody = {
        User: CONSTS.USER,
        Book: book,
      };
      this.tryReturnBook(reqBody);
    },
  },
  data() {
    return {
      search: "",
      calories: "",
      coins: [],
    };
  },
};
</script>
