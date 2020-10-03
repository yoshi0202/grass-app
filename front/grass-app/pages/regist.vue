<template>
  <v-container text-center>
    <v-card flat tile class="text-h1 d-flex justify-center">YourPages</v-card>
    <v-container>
      <v-row>
        <v-col md="4" offset-md="4">
          <v-container m-0 p-0 d-flex justify-space-around>
            <v-avatar class="mr-5" size="100">
              <img
                :src="'https://github.com/' + ownData.login + '.png'"
                :alt="ownData.login"
              />
            </v-avatar>
            <v-container m-0 p-0>
              <p>GitHubID: {{ ownData.login }}</p>
              <p>
                YourPage:
                <a :href="'https://github.com/' + ownData.login"
                  >https://github.com/{{ ownData.login }}</a
                >
              </p>
            </v-container>
          </v-container>
          <v-container>
            <v-card flat tile outlined>
              <v-container>
                <v-container class="text-h5">
                  GitHubIDを入力
                  <v-container>
                    <v-container class="text-subtitle-1">
                      <p>入力されたGitHubユーザの草が</p>
                      <p>Timelineに表示されるようになります。</p>
                    </v-container>
                    <v-text-field
                      label="GitHubID"
                      single-line
                      outlined
                      v-model="gitHubID"
                    ></v-text-field>
                    <v-btn @click="regist()" outlined>Register</v-btn>
                  </v-container>
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
  data: function () {
    return {
      ip: "aiueo",
      gitHubID: "",
      ownData: "aiueo",
    };
  },
  mounted: async function () {
    this.getOwn();
  },
  methods: {
    async regist() {
      console.log(this.gitHubID);
      const registResult = await this.$axios.$post(
        "http://localhost:8080/timeline/regist",
        {
          gitHubID: this.gitHubID,
        },
        {
          withCredentials: true,
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
    },
    async getOwn() {
      const ownData = await this.$axios.$get("http://localhost:8080/user", {
        withCredentials: true,
      });
      this.ownData = ownData;
      if (ownData === "redirect") {
        this.$router.push("/login");
        return;
      }
    },
  },
};
</script>
