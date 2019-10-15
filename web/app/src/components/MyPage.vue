<template>
  <v-app id="inspire">
    <SideMenu/>
    <v-app-bar
      app
      color="indigo"
      dark
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container
        fluid
        fill-height
      >
        <v-layout
          align-center
          justify-center
        >
          <v-flex text-xs-center>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <BottomNavigator/>
    <v-footer
      color="indigo"
      app
    >
      <span class="white--text">&copy; 2019</span>
    </v-footer>
  </v-app>
</template>
<script>
import BottomNavigator from './BottomNav';
import SideMenu from './SideMenu';
export default {
  data() {
    return {
      drawer: null
    }
  },
  components: { SideMenu, BottomNavigator },
  methods: {
  },
}
</script>