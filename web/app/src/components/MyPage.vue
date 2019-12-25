<template>
  <v-app id="inspire">
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
            {{ this.midText }}
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer
      color="indigo"
      app
    >
      <span class="white--text">&copy; 2019</span>
    </v-footer>
  </v-app>
</template>
<script>
import client from '../client'
export default {
  data() {
    return {
      midText: "Fetching..",
      drawer: null
    }
  },
  components: { },
  created () {
    this.getHelloWorld();
  },
  watch: {
    '$route': 'getHelloWorld'
  },
  methods: {
    getHelloWorld() {
      client.getHelloWorld().then((resp) => {
        if(resp.status) {
          this.midText = resp.message;
        } else {
          this.midText = "Failed";
        }
      })
    }
  },
}
</script>