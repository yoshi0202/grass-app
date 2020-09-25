<template>
  <v-container>
    <v-card flat tile class="text-h1 d-flex justify-center">Timeline</v-card>
    <v-container text-center>
      <v-row>
        <v-col sm="10" offset-sm="1">
          <v-container v-for="t of timelineArr" :key="t.id" text-center>
            <no-ssr placeholder="Loading...">
              <v-avatar>
                <img
                  :src="'https://github.com/' + t.githubId + '.png'"
                  :alt="t.githubId"
                />
              </v-avatar>
              <a :href="'https://github.com/' + t.githubId" target="_blank">
                {{ t.githubId }}
              </a>
              <v-btn text small @click="destroy(t.id)">delete</v-btn>
              <v-container text-center>
                <github-lawn
                  :data="t.count"
                  :unit="'contributions'"
                ></github-lawn>
              </v-container>
            </no-ssr>
          </v-container>
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<script>
import GithubLawn from "vue-github-lawn";

export default {
  components: {
    GithubLawn
  },
  data: function() {
    return {
      timelineArr: []
    };
  },
  mounted: function() {
    this.initTimeline();
  },
  methods: {
    async destroy(grassId) {
      await this.$axios.$put(
        "http://localhost:8080/timeline/delete",
        {
          grassID: grassId
        },
        { withCredentials: true }
      );
      this.initTimeline();
    },
    async initTimeline() {
      const timeline = await this.$axios.$get(
        "http://localhost:8080/timeline",
        {
          withCredentials: true
        }
      );
      if (timeline === "redirect") {
        this.$router.push("/login");
        return;
      }
      const arr = [];
      for (const t of timeline) {
        let obj = {};
        obj.githubId = t.githubId;
        obj.id = t.ID;
        obj.count = [];
        for (const c of t.countDate) {
          obj.count.push(c.count);
        }
        arr.push(obj);
      }
      this.timelineArr = arr;
    }
  }
};
</script>

<style scoped>
.github-lawn {
  display: inline-block;
}
</style>
