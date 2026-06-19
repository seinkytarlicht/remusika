<script setup lang="ts">
const isUseDrawer = useLocalStorage("remusika_use_drawer", true);
</script>

<template>
  <UContainer class="h-full max-w-180" v-if="!isUseDrawer">
    <MusicList />
  </UContainer>
</template>
