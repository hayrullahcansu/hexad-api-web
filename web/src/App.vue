<template>
  <v-app id="inspire">
    <v-navigation-drawer class="deep-purple accent-4" dark app v-model="drawer">
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6"> Library </v-list-item-title>
          <v-list-item-subtitle> Management </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list>
        <v-list-item v-for="item in items" :key="item.title" :to="item.to" link>
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <div class="pa-2">
          <span>User: </span> <v-chip class="ma-1"> Test123 </v-chip>
        </div>
      </template>
    </v-navigation-drawer>
    <v-app-bar app color="deep-purple accent-4">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Library Management</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <!--  -->
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
export default {
  mounted() {
    // this.$router.push({ path: `/books` })
    // const session = auth.getSession();
    // if (session && session.jwt) {
    //   this.setSession(session);
    //   this.tryGetUserInfo();
    // }
  },
  data() {
    return {
      drawer: true,
      group: null,
      items: [
        { title: "Books", icon: "mdi-home-variant", to: "/books" },
        {
          title: "My Borrowed List",
          icon: "mdi-star-circle",
          to: "/borrowed-list",
        },
        { title: "About", icon: "mdi-help-box", to: "/about" },
      ],
    };
  },
};
</script>
