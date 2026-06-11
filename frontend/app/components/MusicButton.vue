<script setup lang="ts">
import { getColorSync, getPaletteSync, type BrowserSource } from "colorthief";
import type { Music } from "~/types/music";

const route = useRoute();
const imageRef = ref<HTMLImageElement>();
const colorLightImg = ref<string>();
const colorDarkImg = ref<string>();
const colorButton = ref<string>();
const colorMode = useColorMode();

type PageProps = {
  music: Music;
};

const { music: m } = defineProps<PageProps>();

watch(imageRef, (img) => {
  if (!img) return;

  img.addEventListener("load", () => {
    const palette = getPaletteSync(img, { colorCount: 5 });

    if (!palette) return;

    colorDarkImg.value = (
      palette.find((color) => color.isDark) ?? undefined
    )?.hex();
    colorLightImg.value = (
      palette.find((color) => color.isLight) ?? undefined
    )?.hex();
  });
});

watch(
  [() => colorMode.preference, colorLightImg, colorDarkImg],
  ([preference]) => {
    if (preference == "dark") {
      colorButton.value = colorLightImg.value;
    } else {
      colorButton.value = colorDarkImg.value;
    }
  },
);
</script>

<template>
  <NuxtLink :to="`/play/${m.uuid}`" :key="m.uuid">
    <div
      class="p-2 rounded-lg flex gap-3"
      :style="{
        '--ui-primary': colorButton,
      }"
      :class="
        route.params.uuid === m.uuid
          ? 'bg-primary/20 text-primary font-bold'
          : 'hover:bg-primary/20 text-default hover:text-primary'
      "
    >
      <div
        class="h-14 aspect-square flex justify-center items-center bg-elevated rounded-md overflow-hidden"
      >
        <img
          v-if="m.image_url"
          :src="m.image_url"
          class="object-cover"
          ref="imageRef"
        />
        <UIcon v-else name="i-lucide-music-2" class="size-4" />
      </div>

      <div class="flex flex-col justify-center flex-1 min-w-0">
        <h6 class="truncate">{{ m.title }}</h6>
        <p class="text-sm text-muted truncate">
          {{ m.artist }} - {{ m.album }}
        </p>
      </div>
    </div>
  </NuxtLink>
</template>
