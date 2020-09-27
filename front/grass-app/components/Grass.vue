<template>
  <v-container text-center>
    <v-row>
      <v-col cols="6" offset="3">
        <v-container my-0 py-0 d-flex justify-space-around>
          <v-container text-left my-0 py-0>
            <v-avatar class="mr-5">
              <img
                :src="'https://github.com/' + data.githubId + '.png'"
                :alt="data.githubId"
              />
            </v-avatar>
            <a :href="'https://github.com/' + data.githubId" target="_blank">
              {{ data.githubId }}
            </a>
          </v-container>
          <v-container v-if="!isOwn" text-right>
            <v-icon @click="destroy(data.id)">{{
              mdiCloseCircleOutline
            }}</v-icon>
          </v-container>
          <v-container v-else> </v-container>
        </v-container>
      </v-col>
    </v-row>
    <v-container text-center my-0 py-0>
      <github-lawn :data="data.count" :unit="'contributions'"></github-lawn>
    </v-container>
  </v-container>
</template>

<script>
import GithubLawn from "vue-github-lawn";
import { mdiGithub, mdiCloseCircleOutline } from "@mdi/js";
export default {
  props: ["data", "isOwn"],
  components: {
    GithubLawn
  },
  data: function() {
    return {
      mdiGithub: mdiGithub,
      mdiCloseCircleOutline: mdiCloseCircleOutline
    };
  }
};
</script>

<style scoped>
.github-lawn {
  display: inline-block;
}
</style>
