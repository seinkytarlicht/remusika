<script setup lang="ts">
const musicStore = useMusicStore();
</script>

<template>
  <div class="flex flex-col gap-4 flex-1 overflow-y-auto h-full">
    <UInput
      trailing-icon="i-ph-magnifying-glass"
      size="lg"
      variant="outline"
      placeholder="Search by Title, Artist, or Album"
      v-model="musicStore.search"
    >
      <template v-if="musicStore.search?.length" #trailing>
        <UButton
          color="neutral"
          variant="link"
          size="sm"
          icon="i-ph-x-circle"
          aria-label="Clear input"
          @click="musicStore.search = ''"
        />
      </template>
    </UInput>

    <div
      class="flex flex-col gap-2 overflow-auto scrollbar-thin scrollbar-track-transparent scrollbar-thumb-slate-800"
    >
      <template v-if="!musicStore.loading">
        <MusicButton
          :music="m"
          v-for="m in musicStore.music"
          :key="m.uuid"
        ></MusicButton>
      </template>
    </div>
  </div>
</template>
