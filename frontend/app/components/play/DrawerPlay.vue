<script setup lang="ts">
import { formatTime } from "~/helper/time";

const musicStore = useMusicStore();
const playerStore = usePlayerStore();
const isUseDrawer = useLocalStorage("remusika_use_drawer", true);

const drawerEl = ref();
const dragY = ref(0);
const isDragging = ref(false);

const CLOSE_THRESHOLD = 37 / 100;

const { distanceY } = usePointerSwipe(drawerEl, {
  threshold: 0,
  onSwipeStart() {
    isDragging.value = true;
  },
  onSwipe() {
    const delta = -distanceY.value;
    dragY.value = Math.max(0, delta);
  },
  onSwipeEnd() {
    isDragging.value = false;
    if (dragY.value > drawerEl.value.offsetHeight * CLOSE_THRESHOLD) {
      isUseDrawer.value = false;
    }
    dragY.value = 0;
  },
});

defineShortcuts({
  escape: {
    handler: () => (isUseDrawer.value = false),
  },
});
</script>

<template>
  <div
    class="sheet absolute flex w-full h-full bottom-0 left-0 z-50 bg-default"
    :class="{
      'no-transition': isDragging,
    }"
    :style="{ '--drag-y': !isUseDrawer ? '100%' : `${dragY}px` }"
  >
    <UButton
      icon="i-ph-x-bold"
      size="xl"
      color="primary"
      variant="soft"
      @click="isUseDrawer = false"
      class="absolute top-4 right-4 z-10"
    />

    <div class="w-[30vw] px-4 py-5">
      <MusicList />
    </div>

    <main
      class="flex flex-col justify-center items-center flex-1 h-full gap-8 relative"
    >
      <!-- Background -->
      <div
        ref="drawerEl"
        class="absolute top-0 left-0 w-full h-full overflow-hidden light:bg-primary-100 dark:bg-primary-900"
      >
        <template v-if="musicStore.currentMusic?.image_url">
          <img
            :src="musicStore.currentMusic.image_url"
            class="w-full h-full object-cover scale-125 relative select-none pointer-events-none blur-xl"
            draggable="false"
            loading="lazy"
          />
          <div class="absolute w-full h-full top-0 left-0 bg-black/50"></div>
        </template>
      </div>

      <!-- Image Thumbnail -->
      <div
        class="relative h-[50%] aspect-square flex justify-center items-center dark:bg-elevated/50 light:bg-elevated rounded-3xl overflow-hidden shadow-primary/20 shadow-2xl pointer-events-none select-none"
      >
        <img
          v-if="musicStore.currentMusic?.image_url"
          :src="musicStore.currentMusic.image_url"
          :alt="musicStore.currentMusic.title"
          class="w-full h-full object-cover select-none pointer-events-none"
          draggable="false"
          loading="lazy"
        />
        <UIcon v-else name="i-ph-music-note-fill" class="size-[20%]" />
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
              <UIcon name="i-ph-skip-back-fill" class="size-[90%]" />
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
              @click="() => playerStore.playMusic()"
            >
              <UIcon
                :name="
                  !playerStore.isPlaying ? 'i-ph-play-fill' : 'i-ph-pause-fill'
                "
                class="size-[70%]"
              />
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
              <UIcon name="i-ph-skip-forward-fill" class="size-[90%]" />
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
                    ? 'i-ph-repeat-bold'
                    : 'i-ph-repeat-once-bold'
                "
                class="size-full"
              />
            </UButton>
          </UTooltip>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
.sheet {
  transform: translateY(var(--drag-y, 0px));
  transition: transform 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}
.no-transition {
  transition: none !important;
}
</style>
