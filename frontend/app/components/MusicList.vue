<script setup lang="ts">
type PageProps = {
  showedPlaylist?: boolean;
};

const { showedPlaylist = false } = defineProps<PageProps>();

const musicStore = useMusicStore();
const playlistStore = usePlaylistStore();
const searchCurrent = computed({
  get: () =>
    showedPlaylist ? musicStore.searchShowed : musicStore.searchQueue,
  set: (val) => {
    if (showedPlaylist) {
      musicStore.searchShowed = val;
    } else {
      musicStore.searchQueue = val;
    }
  },
});
</script>

<template>
  <div class="flex flex-col gap-4 flex-1 overflow-y-auto h-full">
    <h3 class="flex gap-2 items-center text-2xl font-semibold">
      <UIcon name="i-ph-folder-fill" />
      {{
        showedPlaylist
          ? playlistStore.showedPlaylistName
          : playlistStore.selectedPlaylistName
      }}
    </h3>

    <UInput
      trailing-icon="i-ph-magnifying-glass"
      size="lg"
      variant="outline"
      :placeholder="`Search in ${showedPlaylist ? playlistStore.showedPlaylistName : playlistStore.selectedPlaylistName} by Title, Artist, or Album`"
      v-model="searchCurrent"
    >
      <template v-if="searchCurrent.length > 0" #trailing>
        <UButton
          color="neutral"
          variant="link"
          size="sm"
          icon="i-ph-x-circle"
          aria-label="Clear input"
          @click="searchCurrent = ''"
        />
      </template>
    </UInput>

    <div
      class="flex flex-col gap-2 overflow-auto scrollbar-thin scrollbar-track-transparent scrollbar-thumb-slate-800"
    >
      <template v-if="showedPlaylist">
        <MusicButton
          :music="m"
          v-for="m in musicStore.showedPlaylist"
          :key="m.uuid"
          :playlist="playlistStore.showedPlaylistName"
        />
      </template>
      <template v-else>
        <MusicButton
          :music="m"
          v-for="m in musicStore.music"
          :key="m.uuid"
          :playlist="playlistStore.selectedPlaylistName"
        />
      </template>
    </div>
  </div>
</template>
