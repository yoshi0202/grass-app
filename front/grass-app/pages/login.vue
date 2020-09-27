<template>
  <v-container>
    <v-card flat tile class="text-h1 d-flex justify-center">Login</v-card>
    <v-container text-center>
      <v-row>
        <v-col sm="4" offset-sm="4">
          <v-container>
            <v-card flat tile outlined>
              <v-container>
                <v-container class="text-h5">
                  SignIn With GitHub
                  <v-container mb-6>
                    <v-layout justify-center>
                      <v-img
                        lazy-src="/GitHub-Mark-120px-plus.png"
                        src="/GitHub-Mark-120px-plus.png"
                        max-height="100"
                        max-width="100"
                      ></v-img>
                    </v-layout>
                  </v-container>
                  <v-btn
                    outlined
                    href="https://github.com/login/oauth/authorize?login=grass-app&client_id=6275811b400983aaba52&redirect_uri=http://localhost:8080/github/callback"
                    >SignIN</v-btn
                  >
                </v-container>
              </v-container>
            </v-card>
          </v-container>
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
