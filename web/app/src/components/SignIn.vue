<template>
  <div id="app">
    <v-app id="inspire">
      <v-content>
        <v-container
          fluid
          fill-height
        >
          <v-layout
            align-center
            justify-center
          >
            <v-flex
              xs12
              sm8
              md4
            >
              <v-card class="elevation-12">
                <v-toolbar
                  color="primary"
                  dark
                  flat
                >
                  <v-toolbar-title>Login form</v-toolbar-title>
                  <v-spacer></v-spacer>
                </v-toolbar>
                <v-card-text>
                  <v-form>
                    <v-text-field
                      label="Login"
                      name="login"
                      prepend-icon="mdi-account"
                      type="text"
                      v-model="userInfo.email">
                    </v-text-field>

                    <v-text-field
                      id="password"
                      label="Password"
                      name="password"
                      prepend-icon="mdi-lock"
                      append-icon="mdi-eye-off"
                      :type="showPassword ? 'text' : 'password'"
                      v-model="userInfo.password"
                      @click:append="toggleShowPassword()">
                    </v-text-field>
                  </v-form>
                </v-card-text>
                <v-card-actions>
                  <v-btn color="primary" @click="$router.push('/signup')">Register</v-btn>
                  <v-spacer></v-spacer>
                  <v-btn color="primary" @click="login()">Login</v-btn>
                </v-card-actions>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
    </v-app>
  </div>
</template>
<script>
  import client from '../client'
  export default {
    props: {
    },
    data() {
      return {
        showPassword: false,
        userInfo: {
          email: "",
          password: ""
        }
      }
    },
    async mounted() {

    },
    methods: {
      toggleShowPassword(){
        this.showPassword = !this.showPassword;
      },
      login(){
        client.login(JSON.stringify(this.userInfo))
        .then((resp) => {
          if (resp.status) {
            this.$auth.setAccessToken(resp.token);
            this.$router.push("/my_page");
          } else {
            this.flash(resp.message, 'error');
          }
        });
      },
    },
}
</script>