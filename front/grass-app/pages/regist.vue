<template>
  <v-container>
    <v-card flat tile class="text-h1 d-flex justify-center">Register</v-card>
    <v-row>
      <v-col md="4" offset-md="4">
        <v-container text-center>
          <v-card tile>
            input GitHubID
            <v-container>
              <v-text-field
                label="GitHubID"
                single-line
                outlined
                v-model="gitHubID"
              ></v-text-field>
              <v-btn text small @click="regist()">Register</v-btn>
            </v-container>
          </v-card>
        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: function() {
    return {
      ip: "aiueo",
      gitHubID: ""
    };
  },
  mounted: async function() {},
  methods: {
    async regist() {
      console.log(this.gitHubID);
      const registResult = await this.$axios.$post(
        "http://localhost:8080/timeline/regist",
        {
          gitHubID: this.gitHubID
        },
        {
          withCredentials: true
        }
      );
      if (registResult === "redirect") {
        this.$router.push("/login");
        return;
      }
      if (registResult === "success!") {
        this.$router.push("/timeline");
        return;
      }
    }
  }
};
</script>
