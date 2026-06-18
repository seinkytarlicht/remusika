<script setup lang="ts">
import MusicList from "~/components/MusicList.vue";
import DrawerPlay from "~/components/play/DrawerPlay.vue";
import { formatTime } from "~/helper/time";

const isUseDrawer = useLocalStorage("remusika_use_drawer", true);
const musicStore = useMusicStore();
const playerStore = usePlayerStore();

watch(
  [() => musicStore.loading, () => musicStore.currentMusicError],
  ([loading, err]) => {
    if (!loading && err) {
      navigateTo("/");
    }
  },
);

defineShortcuts({
  shift_alt_n: {
    handler: () => {
      playerStore.prevMusic();
    },
  },
  shift_n: {
    handler: () => {
      playerStore.nextMusic();
    },
  },
  k: {
    handler: () => {
      playerStore.playMusic();
    },
  },
  l: {
    handler: () => {
      playerStore.toggleLoop();
    },
  },
});
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

  <UContainer class="h-full max-w-180" v-if="!isUseDrawer">
    <MusicList />
  </UContainer>

  <DrawerPlay v-model:drawer-open="isUseDrawer" />

  <Teleport to="#footer-main-panel">
    <USlider
      class="z-10"
      v-model="playerStore.currentTime"
      :min="0"
      :max="playerStore.duration"
      @pointerdown="playerStore.startSeek"
      @update:model-value="playerStore.seekMusic"
      @pointerup="playerStore.endSeek"
      :ui="{
        range: 'rounded-none',
        track: 'rounded-none',
      }"
    />
    <div
      v-if="playerStore.audioRef"
      class="py-4 px-4 dark:bg-elevated/50 bg-elevated w-full flex flex-col items-center justify-center gap-2 relative"
    >
      <div class="flex justify-between gap-8 items-center w-full max-w-full">
        <div class="min-w-0 w-full cursor-pointer" @click="isUseDrawer = true">
          <div
            class="p-2 rounded-lg flex gap-3 group"
            v-if="musicStore.currentMusic"
          >
            <div
              class="h-14 aspect-square flex justify-center items-center bg-elevated rounded-md overflow-hidden"
            >
              <img
                v-if="musicStore.currentMusic.image_url"
                :src="musicStore.currentMusic.image_url"
                class="object-cover"
                ref="imageRef"
              />
              <UIcon v-else name="i-lucide-music-2" class="size-4" />
            </div>

            <div class="flex flex-col justify-center flex-1 min-w-0">
              <h6 class="truncate group-hover:underline">
                {{ musicStore.currentMusic.title }}
              </h6>
              <p class="text-sm text-muted truncate">
                {{ musicStore.currentMusic.artist }} -
                {{ musicStore.currentMusic.album }}
              </p>
            </div>
          </div>
        </div>

        <!-- Controls Start -->
        <div class="flex gap-4 justify-between items-center w-max-120">
          <div class="flex justify-between items-center gap-2">
            <p>
              {{ formatTime(playerStore.currentTime) }}
            </p>
            <p>/</p>
            <p>
              {{ formatTime(playerStore.duration) }}
            </p>
          </div>

          <div />

          <!-- Prev -->
          <UTooltip
            text="Prev"
            :kbds="['shift', 'alt', 'n']"
            :delay-duration="0"
            :content="{ side: 'top' }"
          >
            <UButton
              color="primary"
              variant="soft"
              class="rounded-full size-12 flex justify-center"
              @click="playerStore.prevMusic"
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
              variant="soft"
              class="rounded-full size-18 flex justify-center"
              @click="playerStore.playMusic"
            >
              <!-- the inline thing doesn't work somehow  -->
              <UIcon
                v-if="!playerStore.isPlaying"
                name="i-lucide-play"
                class="size-[70%]"
              />
              <UIcon v-else name="i-lucide-pause" class="size-[80%]" />
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
              variant="soft"
              class="rounded-full size-12 flex justify-center"
              @click="playerStore.nextMusic"
            >
              <UIcon name="i-lucide-skip-forward" class="size-[90%]" />
            </UButton>
          </UTooltip>

          <div />

          <!-- Loop -->
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
              @click="playerStore.toggleLoop"
            >
              <UIcon
                :name="
                  !playerStore.isLooping
                    ? 'i-lucide-repeat'
                    : 'i-lucide-repeat-1'
                "
                class="size-full"
              />
            </UButton>
          </UTooltip>
        </div>
        <!-- Controls End -->
      </div>
    </div>
  </Teleport>
</template>
