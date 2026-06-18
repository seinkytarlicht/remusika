<script setup lang="ts">
import MusicButton from "~/components/MusicButton.vue";
import MusicList from "~/components/MusicList.vue";

useHead({
  title: "ReMusika",
});

const route = useRoute();
const { $api } = useNuxtApp();
const toast = useToast();
const playerStore = usePlayerStore();
const musicStore = useMusicStore();
const playlistStore = usePlaylistStore();
const { initSession, updatePlaybackState, setNextTrack, setPrevTrack } =
  useMediaSession();

const isUseDrawer = useLocalStorage("remusika_use_drawer", true);
const audioRef = ref<HTMLAudioElement>();

watch(audioRef, (audio) => {
  if (audio) {
    playerStore.audioRef = audio;
  }
});

watch(
  () => musicStore.currentMusic,
  () => {
    if (musicStore.currentMusic) {
      initSession(musicStore.currentMusic);
      setPrevTrack(playerStore.prevMusic);
      setNextTrack(playerStore.nextMusic);
    }
  },
);

watch(
  () => playerStore.isPlaying,
  () => {
    if (playerStore.isPlaying) {
      updatePlaybackState("playing");
    } else {
      updatePlaybackState("paused");
    }
  },
);

watch(
  [() => musicStore.loading, () => route.params.uuid],
  ([newLoading, newUuid]) => {
    if (!newLoading) {
      playerStore.nowPlaying = musicStore.setCurrentMusic(String(newUuid));
    }
  },
  { immediate: true },
);

watch(
  () => playerStore.isEnded,
  (isEnded) => {
    if (isEnded) playerStore.nextMusic();
  },
);

onMounted(() => {
  musicStore.fetch();
});

async function reloadPlaylist() {
  try {
    await $api("/music/reload-folder", {
      method: "POST",
    });

    toast.add({
      title: "Playlist Reloaded",
    });

    musicStore.fetch();
  } catch (error) {
    console.error(error);
  }
}

async function shutdownSystem() {
  try {
    await $api("/shutdown/off", {
      method: "POST",
    });

    toast.add({
      title: "the App has been stopped",
      description: "You can close this tab",
    });
  } catch (error) {
    console.error(error);
  }
}
</script>

<template>
  <UDashboardGroup storage-key="remusika" storage="local">
    <!-- Playlist Sidebar Start -->
    <UDashboardSidebar
      id="playlist-panel"
      resizable
      :default-size="20"
      :max-size="20"
      :min-size="15"
    >
      <template #resize-handle="{ onMouseDown, onTouchStart, onDoubleClick }">
        <UDashboardResizeHandle
          class="after:absolute after:inset-y-0 after:right-0 after:w-px hover:after:bg-(--ui-border-accented) after:transition"
          @mousedown="onMouseDown"
          @touchstart="onTouchStart"
          @dblclick="onDoubleClick"
        />
      </template>

      <template #header>
        <h1 class="text-3xl font-bold flex gap-2 items-center w-full">
          <UIcon name="i-lucide-list-music" class="size-8" /> Playlist
        </h1>

        <UTooltip text="Reload Playlist">
          <UButton
            icon="i-lucide-refresh-ccw"
            color="primary"
            size="lg"
            variant="outline"
            @click="reloadPlaylist"
          />
        </UTooltip>
      </template>

      <template #default="{ collapsed }">
        <div v-for="pl in playlistStore.playlist">
          <p>{{ pl.name }}</p>
        </div>
      </template>

      <template #footer>
        <Logo />
        <UTooltip text="Shutdown" :delay-duration="0">
          <UButton
            icon="i-lucide-power"
            color="error"
            variant="outline"
            @click="shutdownSystem"
          />
        </UTooltip>
        <ConfigureModal />
      </template>
    </UDashboardSidebar>
    <!-- Playlist Sidebar End -->

    <!-- Music Sidebar Start -->
    <UDashboardSidebar
      id="music-panel"
      :collapsed="!isUseDrawer"
      collapsible
      :resizable="isUseDrawer"
      :collapsed-size="0"
      :min-size="20"
      :max-size="30"
      :default-size="30"
      :ui="{ footer: 'border-t border-default' }"
      class="min-w-0"
    >
      <template #resize-handle="{ onMouseDown, onTouchStart, onDoubleClick }">
        <UDashboardResizeHandle
          v-if="isUseDrawer"
          class="after:absolute after:inset-y-0 after:right-0 after:w-px hover:after:bg-(--ui-border-accented) after:transition"
          @mousedown="onMouseDown"
          @touchstart="onTouchStart"
          @dblclick="onDoubleClick"
        />
      </template>

      <template #default="{ collapsed }">
        <MusicList v-if="!collapsed" />
      </template>
    </UDashboardSidebar>
    <!-- Music Sidebar End -->

    <!-- Main Panel Start -->
    <UDashboardPanel>
      <template #body>
        <audio
          ref="audioRef"
          autoplay
          @play="playerStore.isPlaying = true"
          @pause="playerStore.isPlaying = false"
          :src="musicStore.currentMusic?.audio_url"
        ></audio>

        <slot />
      </template>

      <template #footer>
        <div id="footer-main-panel" />
      </template>
    </UDashboardPanel>
    <!-- Main Panel End -->

    <!-- Metadata Panel Start -->
    <UDashboardSidebar
      id="metadata-panel"
      :resizable="true"
      side="right"
      collapsible
      :collapsed-size="0"
      :min-size="8"
      :max-size="15"
      :default-size="15"
      class="min-w-0"
    >
      <template #resize-handle="{ onMouseDown, onTouchStart, onDoubleClick }">
        <UDashboardResizeHandle
          class="after:absolute after:inset-y-0 after:right-0 after:w-px hover:after:bg-(--ui-border-accented) after:transition"
          @mousedown="onMouseDown"
          @touchstart="onTouchStart"
          @dblclick="onDoubleClick"
        />
      </template>

      <template #default="{ collapsed }">
        <div v-show="!collapsed" class="wrap-break-word" id="metadata"></div>
      </template>
    </UDashboardSidebar>
    <!-- Metadata Panel End -->
  </UDashboardGroup>
</template>
