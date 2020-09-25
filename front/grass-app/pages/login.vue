<template>
  <v-container>
    <v-card flat tile class="text-h1 d-flex justify-center">Login</v-card>
    <v-container text-center>
      <v-row>
        <v-col sm="4" offset-sm="4">
          <v-card
            tile
            href="https://github.com/login/oauth/authorize?login=grass-app&client_id=6275811b400983aaba52&redirect_uri=http://localhost:8080/github/callback"
            >Signed in GitHub</v-card
          >
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<script>
export default {
  methods: {
    clickSignIn() {
      alert("github");
    }
  }
};
</script>
