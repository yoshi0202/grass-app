<template>
  <v-container>
    <v-row>
      <v-col cols="8" offset="2">
        <v-card flat tile class="text-h1 d-flex justify-center"
          >Timeline</v-card
        >
        <v-container text-center>
          <v-container text-center>
            <p class="text-h5">強い人たちのContributionを監視しよう！！</p>
            <v-row>
              <v-col cols="12">
                <v-card flat outlined>
                  <v-container text-h4 mb-0 pb-0
                    >あなたの<span class="text-h3 green--text"
                      >草🌱</span
                    ></v-container
                  >
                  <v-container>
                    <no-ssr placeholder="Loading...">
                      <Grass :data="ownGrass" :isOwn="true" />
                    </no-ssr>
                  </v-container>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-container>
        <v-container text-center>
          <v-row>
            <v-col cols="12">
              <v-card flat outlined>
                <v-container text-h4 mb-0 pb-0
                  >強い人たちの<span class="text-h3 green--text"
                    >草💪</span
                  ></v-container
                >
                <v-container v-for="t of timelineArr" :key="t.id" text-center>
                  <no-ssr placeholder="Loading...">
                    <Grass :data="t" :isOwn="false" />
                  </no-ssr>
                </v-container>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import GithubLawn from "vue-github-lawn";
import Grass from "../components/Grass";

export default {
  components: {
    GithubLawn,
    Grass,
  },
  data: function () {
    return {
      timelineArr: [],
      ownGrass: {
        count: [],
        githubId: "",
        id: "",
      },
    };
  },
  mounted: function () {
    this.initTimeline();
    this.createMyGrass();
  },
  methods: {
    async destroy(grassId) {
      await this.$axios.$put(
        "http://localhost:8080/timeline/delete",
        {
          grassID: grassId,
        },
        { withCredentials: true }
      );
      this.initTimeline();
    },
    async initTimeline() {
      const timeline = await this.$axios.$get(
        "http://localhost:8080/timeline",
        {
          withCredentials: true,
        }
      );
      if (timeline === "redirect") {
        this.$router.push("/login");
        return;
      }
      for (const t of timeline) {
        let obj = {};
        obj.githubId = t.githubId;
        obj.id = t.ID;
        obj.count = [];
        for (const c of t.countDate) {
          obj.count.push(c.count);
        }
        this.timelineArr.push(obj);
      }
    },
    async createMyGrass() {
      const own = await this.$axios.$get("http://localhost:8080/owngrass", {
        withCredentials: true,
      });
      if (own === "redirect") {
        this.$router.push("/login");
        return;
      }
      this.ownGrass.githubId = own.githubId;
      this.ownGrass.id = own.id;
      for (const c of own.countDate) {
        this.ownGrass.count.push(c.count);
      }
    },
  },
};
</script>

