<script setup lang="ts">
import { formatTime } from "~/helper/time";
import type { Music } from "~/types/music";

const route = useRoute();
const musicStore = useMusicStore();

const audioRef = ref<HTMLAudioElement | null>(null);
const duration = ref(0);
const currentTime = ref(0);
const isSeeking = ref(false);
const isPlaying = ref(false);
const isLooping = ref(false);
const nowPlaying = ref<Music>();

watch(audioRef, (audio) => {
  if (!audio) return;

  const onLoadedMetadata = () => {
    duration.value = audio.duration;
  };

  const onTimeUpdate = () => {
    if (!isSeeking.value) currentTime.value = audio.currentTime;
  };
  onTimeUpdate();

  audio.addEventListener("loadedmetadata", onLoadedMetadata);
  audio.addEventListener("ended", nextMusic);
  audio.addEventListener("timeupdate", onTimeUpdate);
});

watch(
  [() => musicStore.loading, () => route.params.uuid],
  ([newLoading, newUuid]) => {
    if (!newLoading && newUuid) {
      nowPlaying.value = musicStore.setCurrentMusic(String(newUuid));
    }
  },
  { immediate: true },
);

defineShortcuts({
  shift_alt_n: {
    handler: () => {
      prevMusic();
    },
  },
  shift_n: {
    handler: () => {
      nextMusic();
    },
  },
  k: {
    handler: () => {
      playMusic();
    },
  },
  l: {
    handler: () => {
      toggleLoop();
    },
  },
});

function prevMusic() {
  musicStore.getPrevSong();
  if (nowPlaying.value?.uuid == musicStore.currentMusic?.uuid) {
    audioRef.value?.play();
  } else {
    navigateTo(`/play/${musicStore.currentMusic?.uuid}`);
  }
}

function nextMusic() {
  musicStore.getNextSong();
  if (nowPlaying.value?.uuid == musicStore.currentMusic?.uuid) {
    audioRef.value?.play();
  } else {
    navigateTo(`/play/${musicStore.currentMusic?.uuid}`);
  }
}

function playMusic() {
  if (!audioRef.value) return;

  if (audioRef.value.paused) {
    audioRef.value.play();
  } else {
    audioRef.value.pause();
  }
}

function startSeek() {
  if (!audioRef.value) return;

  isSeeking.value = true;

  audioRef.value.pause();
}

function seekMusic(value: number | undefined) {
  if (!value) return;

  currentTime.value = Number(value);

  if (audioRef.value) {
    audioRef.value.currentTime = Number(value);
  }
}

function endSeek() {
  isSeeking.value = false;

  if (audioRef.value?.paused) {
    audioRef.value?.play();
  }
}

function toggleLoop() {
  if (audioRef.value) {
    audioRef.value.loop = !audioRef.value.loop;
    isLooping.value = audioRef.value.loop;
  }
}
</script>

<template>
  <Teleport to="#metadata">
    <div class="flex flex-col gap-2 my-6">
      <p class="font-bold">
        Title:
        <span class="font-normal text-muted">
          {{ musicStore.currentMusic?.title }}</span
        >
      </p>
      <p class="font-bold">
        Artist:
        <span class="font-normal text-muted">
          {{ musicStore.currentMusic?.artist }}</span
        >
      </p>
      <p class="font-bold">
        Album:
        <span class="font-normal text-muted">
          {{ musicStore.currentMusic?.album }}</span
        >
      </p>
      <p class="font-bold">
        Path:
        <span class="font-normal text-muted">
          {{ musicStore.currentMusic?.path }}</span
        >
      </p>
    </div>
  </Teleport>

  <div class="flex flex-col justify-center items-center h-full gap-8 relative">
    <UButton
      icon="i-lucide-x"
      size="xl"
      to="/"
      color="primary"
      variant="link"
      class="absolute top-4 right-4 z-10"
    />

    <!-- Background -->
    <div
      v-if="musicStore.currentMusic?.image_url"
      class="absolute top-0 left-0 w-full h-full overflow-hidden"
    >
      <img
        :src="musicStore.currentMusic.image_url"
        class="w-full h-full object-cover scale-125 relative"
      />
      <div
        class="absolute w-full h-full top-0 left-0 backdrop-blur-3xl bg-black/40"
      ></div>
    </div>

    <!-- Image Thumbnail -->
    <div
      class="relative h-[50%] aspect-square flex justify-center items-center bg-elevated/40 rounded-3xl overflow-hidden shadow-white/10 shadow-2xl"
    >
      <img
        v-if="musicStore.currentMusic?.image_url"
        :src="musicStore.currentMusic.image_url"
        alt=""
        class="w-full h-full object-cover"
      />
      <UIcon v-else name="i-lucide-music-2" class="size-[20%]" />
    </div>

    <audio
      :src="nowPlaying?.audio_url"
      ref="audioRef"
      autoplay
      @play="isPlaying = true"
      @pause="isPlaying = false"
    ></audio>

    <div
      v-if="audioRef"
      class="max-w-110 py-6 px-4 dark:bg-elevated/50 bg-elevated w-full flex flex-col items-center justify-center gap-2 rounded-4xl relative"
    >
      <USlider
        v-model="currentTime"
        :min="0"
        :max="duration"
        @pointerdown="startSeek"
        @update:model-value="seekMusic"
        @pointerup="endSeek"
      />
      <div class="flex justify-between items-center w-full">
        <p>
          {{ formatTime(currentTime) }}
        </p>
        <p>
          {{ formatTime(duration) }}
        </p>
      </div>
      <div class="flex gap-8 justify-between items-center w-full">
        <div class="size-12"></div>

        <!-- Prev -->
        <UTooltip
          text="Prev"
          :kbds="['shift', 'alt', 'n']"
          :delay-duration="0"
          :content="{ side: 'top' }"
        >
          <UButton
            color="primary"
            variant="subtle"
            class="rounded-full size-12 flex justify-center"
            @click="prevMusic"
          >
            <UIcon name="i-lucide-skip-back" class="size-[90%]" />
          </UButton>
        </UTooltip>

        <!-- Play -->
        <UTooltip
          text="Play"
          :kbds="['K']"
          :delay-duration="0"
          :content="{ side: 'top' }"
        >
          <UButton
            color="primary"
            variant="subtle"
            class="rounded-full size-20 flex justify-center"
            @click="playMusic"
          >
            <!-- the inline thing doesn't work somehow  -->
            <UIcon v-if="!isPlaying" name="i-lucide-play" class="size-[70%]" />
            <UIcon v-else name="i-lucide-pause" class="size-[70%]" />
          </UButton>
        </UTooltip>

        <!-- Next Song -->
        <UTooltip
          text="Next"
          :kbds="['shift', 'n']"
          :delay-duration="0"
          :content="{ side: 'top' }"
        >
          <UButton
            color="primary"
            variant="subtle"
            class="rounded-full size-12 flex justify-center"
            @click="nextMusic"
          >
            <UIcon name="i-lucide-skip-forward" class="size-[90%]" />
          </UButton>
        </UTooltip>

        <UTooltip
          text="Loop"
          :kbds="['L']"
          :delay-duration="0"
          :content="{ side: 'top' }"
        >
          <UButton
            color="primary"
            variant="ghost"
            class="rounded-full size-12 flex justify-center"
            @click="toggleLoop"
          >
            <UIcon
              :name="!isLooping ? 'i-lucide-repeat' : 'i-lucide-repeat-1'"
              class="size-full"
            />
          </UButton>
        </UTooltip>
      </div>
    </div>
  </div>
</template>
