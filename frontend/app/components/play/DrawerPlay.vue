<script setup lang="ts">
import { formatTime } from "~/helper/time";

const musicStore = useMusicStore();
const playerStore = usePlayerStore();
const isDrawerOpen = defineModel<boolean>("drawerOpen");
</script>

<template>
  <UDrawer
    v-model:open="isDrawerOpen"
    :handle="false"
    :modal="false"
    :dismissible="false"
    :portal="false"
    :overlay="false"
    :ui="{
      content: 'absolute inset-0 h-full max-h-full my-0 z-50',
      overlay: 'absolute inset-0 bg-black/50 z-50',
    }"
  >
    <template #content>
      <UButton
        icon="i-lucide-x"
        size="xl"
        color="primary"
        variant="soft"
        @click="isDrawerOpen = false"
        class="absolute top-4 right-4 z-10"
      />

      <main
        class="flex flex-col justify-center items-center h-full gap-8 relative"
      >
        <!-- Background -->
        <div
          v-if="musicStore.currentMusic?.image_url"
          class="absolute top-0 left-0 w-full h-full overflow-hidden"
        >
          <img
            :src="musicStore.currentMusic.image_url"
            class="w-full h-full object-cover scale-125 relative select-none pointer-events-none"
          />
          <div
            class="absolute w-full h-full top-0 left-0 backdrop-blur-3xl bg-black/40"
          ></div>
        </div>

        <!-- Image Thumbnail -->
        <div
          class="relative h-[50%] aspect-square flex justify-center items-center dark:bg-elevated/50 light:bg-elevated rounded-3xl overflow-hidden shadow-white/10 shadow-2xl"
        >
          <img
            v-if="musicStore.currentMusic?.image_url"
            :src="musicStore.currentMusic.image_url"
            alt=""
            class="w-full h-full object-cover select-none pointer-events-none"
          />
          <UIcon v-else name="i-lucide-music-2" class="size-[20%]" />
        </div>

        <div
          v-if="playerStore.audioRef"
          class="max-w-110 py-6 px-4 dark:bg-elevated/50 bg-elevated w-full flex flex-col items-center justify-center gap-2 rounded-4xl relative"
        >
          <USlider
            v-model="playerStore.currentTime"
            :min="0"
            :max="playerStore.duration"
            @pointerdown="playerStore.startSeek"
            @update:model-value="playerStore.seekMusic"
            @pointerup="playerStore.endSeek"
          />
          <div class="flex justify-between items-center w-full">
            <p>
              {{ formatTime(playerStore.currentTime) }}
            </p>
            <p>
              {{ formatTime(playerStore.duration) }}
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
                class="rounded-full size-20 flex justify-center"
                @click="playerStore.playMusic"
              >
                <!-- the inline thing doesn't work somehow  -->
                <UIcon
                  v-if="!playerStore.isPlaying"
                  name="i-lucide-play"
                  class="size-[70%]"
                />
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
                variant="soft"
                class="rounded-full size-12 flex justify-center"
                @click="playerStore.nextMusic"
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
        </div>
      </main>
    </template>
  </UDrawer>
</template>
