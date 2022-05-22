<template>
  <v-container class="grey lighten-5" fluid>
    <!-- Stack the columns on mobile by making one full-width and the other half-width     class="pa-xs-0 mx-lg-auto" -->
    <h1 class="text-center">Books In The Library</h1>
    <v-row>
      <v-col cols="12" md="12">
        <v-card>
          <v-list v-if="books.length > 0">
            <v-list-item two-line v-for="book in books" :key="book.Name">
              <v-list-item-content>
                <v-list-item-title
                  >{{ book.Name }}
                  <div class="pa-1">
                    <v-btn block @click="borrow(book.Name)"> Borrow </v-btn>
                  </div></v-list-item-title
                >
                <v-list-item-subtitle
                  >Quantity: {{ book.Quantity }}</v-list-item-subtitle
                >
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <div v-else>
            <h4 class="text-center">
              There is no book left to borrow in the library..
            </h4>
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
  name: "BookListing",
  created() {
    this.tryGetBooks();
  },
  computed: {
    ...mapGetters("libraryStore", ["books", "booksLoading"]),
  },
  methods: {
    ...mapActions("libraryStore", ["tryGetBooks", "tryBorrowBook"]),
    filterOnlyCapsText(value, search, item) {
      return (
        value != null &&
        search != null &&
        typeof value === "string" &&
        value.toString().toLocaleUpperCase().indexOf(search) !== -1
      );
    },
    borrow(book) {
      const reqBody = {
        User: CONSTS.USER,
        Book: book,
      };
      this.tryBorrowBook(reqBody);
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
