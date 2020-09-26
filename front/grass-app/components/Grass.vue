<template>
  <div>
    <v-avatar>
      <img
        :src="'https://github.com/' + data.githubId + '.png'"
        :alt="data.githubId"
      />
    </v-avatar>
    <a :href="'https://github.com/' + data.githubId" target="_blank">
      {{ data.githubId }}
    </a>
    <v-btn v-if="!isOwn" text small @click="destroy(data.id)" outlined
      >delete</v-btn
    >
    <v-container text-center>
      <github-lawn :data="data.count" :unit="'contributions'"></github-lawn>
    </v-container>
  </div>
</template>

<script>
import GithubLawn from "vue-github-lawn";
export default {
  props: ["data", "isOwn"],
  components: {
    GithubLawn,
  },
};
</script>

<style scoped>
.github-lawn {
  display: inline-block;
}
</style>