<template>
  <div>
    <v-app>
      <v-content fluid full-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card>
              <v-toolbar flat dark color="green">
                <v-toolbar-title>Registration</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field
                    label="User name"
                    name="name"
                    type="text"
                    v-model="userInfo.name">
                  </v-text-field>
                  <v-text-field
                    label="Email"
                    name="email"
                    type="text"
                    v-model="userInfo.email"
                    rule="email">
                  </v-text-field>
                  <v-text-field
                    label="Password"
                    name="password"
                    :type="showPassword ? 'text' : 'password'"
                    v-model="userInfo.password"
                    @click:append="toggleShowPassword()">
                  </v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="registerUser()">Add me</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-content>
    </v-app>
  </div>
</template>
<script>
import client from "../client"

export default {
  data(){
    return {
      showPassword: false,
      userInfo: {
        name: "",
        email: "",
        password: "",
      }
    }
  },
  methods: {
    toggleShowPassword(){
      this.showPassword = !this.showPassword;
    },
    registerUser(){
      client.request(client.HTTP_POST, "/register", JSON.stringify(this.userInfo));
    }
  }
}
</script>