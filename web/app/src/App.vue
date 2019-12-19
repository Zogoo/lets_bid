<template>
  <div id="app">
    <flash-message class="flashMessages"></flash-message>
    <v-app>
      <router-view></router-view>
    </v-app>
  </div>
</template>

<script>

export default {
  name: 'App',
  data: () => ({
    //
  }),
};
</script>

<style>
  .flashMessages {
    text-align: center;
  }
  .flashMessages button {
    margin: 10px;
    font-size: 20px;
  }
</style>