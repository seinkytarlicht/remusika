<script setup lang="ts">
import MusicButton from "~/components/MusicButton.vue";

useHead({
  title: "ReMusika",
});

const { $api } = useNuxtApp();
const toast = useToast();
const playerStore = usePlayerStore();
const musicStore = useMusicStore();
const { initSession, updatePlaybackState } = useMediaSession();

const audioRef = ref<HTMLAudioElement>();

watch(audioRef, (audio) => {
  if (audio) {
    playerStore.audioRef = audio;
  }
});

watch(
  () => musicStore.currentMusic,
  () => {
    if (musicStore.currentMusic) initSession(musicStore.currentMusic);
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
    <UDashboardSidebar
      resizable
      :min-size="20"
      :max-size="30"
      :default-size="30"
      :ui="{ footer: 'border-t border-default' }"
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

      <UInput
        trailing-icon="i-lucide-search"
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
            icon="i-lucide-circle-x"
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

    <UDashboardSidebar
      resizable
      collapsible
      :collapsed-size="0"
      :min-size="15"
      :max-size="20"
      :default-size="20"
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
        <div v-show="!collapsed" id="metadata"></div>
      </template>
    </UDashboardSidebar>

    <UDashboardPanel>
      <audio
        ref="audioRef"
        autoplay
        @play="playerStore.isPlaying = true"
        @pause="playerStore.isPlaying = false"
        :src="musicStore.currentMusic?.audio_url"
      ></audio>
      <slot />
    </UDashboardPanel>
  </UDashboardGroup>
</template>
