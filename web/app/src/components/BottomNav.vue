<template>
  <v-bottom-navigation
    :value="activeBtn"
    color="deep-purple accent-4"
  >
    <v-btn @click="activateBtn(0); $emit('changeTab', $event, 'history')">
      <span>My page</span>
      <v-icon large color="green darken-2">mdi-home</v-icon>
    </v-btn>

    <v-btn @click="activateBtn(1);  $emit('changeTab', $event, '')">
      <span>Add item</span>
      <v-icon large color="blue darken-2">mdi-arrow-up-bold-box-outline</v-icon>
    </v-btn>

    <v-btn @click="activateBtn(2);  $emit('changeTab', $event, 'history')">
      <span>Message</span>
      <v-icon large color="purple darken-2">mdi-email</v-icon>
    </v-btn>
  </v-bottom-navigation>
</template>
<script>
export default {
  data() {
    return {
      activeBtn: 1,
    }
  },
  methods: {
    activateBtn(number) {
      this.activeBtn = number;
    }
  },
}
</script>