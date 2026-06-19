<script setup lang="ts">
// import { getColorSync, getPaletteSync, type BrowserSource } from "colorthief";
import type { DropdownMenuItem } from "@nuxt/ui";
import type { Music } from "~/types/music";

type PageProps = {
  music: Music;
};

const { $api } = useNuxtApp();
const { music: m } = defineProps<PageProps>();
const route = useRoute();
const playlistStore = usePlaylistStore();

const playlistMenu = computed<DropdownMenuItem[]>(() => {
  const playlists = playlistStore.playlist;

  if (!playlists) return [];

  return [
    {
      label: "Add to playlist",
      icon: "i-ph-folder-fill",
      filter: { placeholder: "Search playlist..." },
      children: playlists.map((p) => ({
        onSelect(e: Event) {
          e.preventDefault();
          const inPlaylist = playlistStore.isMusicInPlaylist(m.uuid, p.name);
          if (inPlaylist) {
            removeMusicFromPlaylist(inPlaylist.id);
          } else {
            addMusicToPlaylist(m.uuid, p.id);
          }
        },
        label: p.name,
        slot: playlistStore.isMusicInPlaylist(m.uuid, p.name)
          ? "in-playlist"
          : "not-in-playlist",
      })),
    },
  ];
});

async function addMusicToPlaylist(uuid: string, playlist_id: number) {
  try {
    await $api("/playlist/add-music", {
      method: "POST",
      body: {
        music_id: uuid,
        playlist_id,
      },
    });

    playlistStore.fetch();
  } catch (e) {
    console.error(e);
  }
}

async function removeMusicFromPlaylist(playlist_item_id: number) {
  try {
    await $api("/playlist/remove-music", {
      method: "POST",
      body: {
        id: playlist_item_id,
      },
    });

    playlistStore.fetch();
  } catch (e) {
    console.error(e);
  }
}
// const imageRef = ref<HTMLImageElement>();
// const colorLightImg = ref<string>();
// const colorDarkImg = ref<string>();
// const colorButton = ref<string>();
// const colorMode = useColorMode();

// watch(imageRef, (img) => {
//   if (!img) return;

//   img.addEventListener("load", () => {
//     const palette = getPaletteSync(img, { colorCount: 5 });

//     if (!palette) return;

//     colorDarkImg.value = (
//       palette.find((color) => color.isDark) ?? undefined
//     )?.hex();
//     colorLightImg.value = (
//       palette.find((color) => color.isLight) ?? undefined
//     )?.hex();
//   });
// });

// watch(
//   [() => colorMode.preference, colorLightImg, colorDarkImg],
//   ([preference]) => {
//     if (preference == "dark") {
//       colorButton.value = colorLightImg.value;
//     } else {
//       colorButton.value = colorDarkImg.value;
//     }
//   },
// );
</script>

<template>
  <NuxtLink
    :to="{
      name: 'play-stream',
      params: {
        uuid: m.uuid,
      },
      query: route.query,
    }"
    :key="m.uuid"
  >
    <div
      class="p-2 rounded-lg flex items-center gap-3"
      :class="
        route.params.uuid === m.uuid
          ? 'bg-primary/20 text-primary font-bold'
          : 'hover:bg-primary/20 text-default hover:text-primary'
      "
    >
      <div
        class="h-14 aspect-square flex justify-center items-center bg-elevated rounded-md overflow-hidden"
      >
        <img v-if="m.image_url" :src="m.image_url" class="object-cover" />
        <UIcon v-else name="i-ph-music-note-fill" class="size-4" />
      </div>

      <div class="flex flex-col justify-center flex-1 min-w-0">
        <h6 class="truncate">{{ m.title }}</h6>
        <p class="text-sm text-muted truncate">
          {{ m.artist }} - {{ m.album }}
        </p>
      </div>

      <UDropdownMenu
        :modal="false"
        :items="playlistMenu"
        :content="{
          align: 'start',
        }"
        :ui="{
          content: 'min-w-48 max-h-54',
        }"
      >
        <UButton
          class="rounded-full"
          icon="i-ph-dots-three-outline-vertical-fill"
          color="neutral"
          variant="ghost"
        />

        <template #not-in-playlist-trailing>
          <UIcon name="i-ph-bookmark-simple-bold" class="shrink-0 size-5" />
        </template>
        <template #in-playlist-trailing>
          <UIcon name="i-ph-bookmark-simple-fill" class="shrink-0 size-5" />
        </template>
      </UDropdownMenu>
    </div>
  </NuxtLink>
</template>
