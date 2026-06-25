<script setup lang="ts">
import type { DropdownMenuItem } from "@nuxt/ui";
import type { Music } from "~/types/music";

type PageProps = {
  music: Music;
  playlist_id: number;
};

const { music: m, playlist_id } = defineProps<PageProps>();

const { $api } = useNuxtApp();
const route = useRoute();
const playlistStore = usePlaylistStore();
const { playMusic } = usePlayerStore();
const isUseDrawer = useLocalStorage("remusika_use_drawer", true);

const dropdownModel = ref(false);

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
          const inPlaylist = playlistStore.isMusicInPlaylist(m.uuid, p.id);
          if (inPlaylist) {
            removeMusicFromPlaylist(inPlaylist.id);
          } else {
            addMusicToPlaylist(m.uuid, p.id);
          }
        },
        label: p.name,
        slot: playlistStore.isMusicInPlaylist(m.uuid, p.id)
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
</script>

<template>
  <div
    class="p-2 rounded-lg flex items-center gap-3 relative"
    v-on:contextmenu="
      (e) => {
        e.preventDefault();
        dropdownModel = true;
      }
    "
    :class="
      route.query['m'] == m.uuid &&
      Number(route.query['playlist']) === playlist_id
        ? 'bg-primary/20 text-primary font-bold'
        : 'hover:bg-primary/20 text-default hover:text-primary'
    "
  >
    <NuxtLink
      :to="{
        name: 'play',
        query: {
          m: m.uuid,
          playlist: playlist_id,
          showlist: route.query['showlist'],
        },
      }"
      @click="
        () => {
          isUseDrawer = true;
          playMusic(true);
        }
      "
      class="absolute top-0 bottom-0 right-0 left-0"
    />

    <div
      class="h-14 aspect-square flex justify-center items-center bg-elevated rounded-md overflow-hidden"
    >
      <img
        v-if="m.image_url"
        :src="m.image_url"
        class="object-cover"
        loading="lazy"
      />
      <UIcon v-else name="i-ph-music-note-fill" class="size-4" />
    </div>

    <div class="flex flex-col justify-center flex-1 min-w-0">
      <template v-if="!m.title || !m.album || !m.artist">
        <h6 class="text-error">Failed to get music details</h6>
        <p class="text-sm text-muted">
          Either deleted or moved, you can remove it from playlist
        </p>
      </template>
      <template v-else>
        <h6 class="truncate">{{ m.title }}</h6>
        <p class="text-sm text-muted truncate">
          {{ m.artist }} - {{ m.album }}
        </p>
      </template>
    </div>

    <UDropdownMenu
      :modal="false"
      v-model:open="dropdownModel"
      :items="playlistMenu"
      :content="{
        align: 'start',
      }"
      :ui="{
        content: 'min-w-48 max-h-54',
      }"
    >
      <UButton
        class="rounded-full relative"
        icon="i-ph-dots-three-outline-vertical-fill"
        color="neutral"
        variant="ghost"
        size="lg"
      />

      <template #not-in-playlist-trailing>
        <UIcon name="i-ph-bookmark-simple-bold" class="shrink-0 size-5" />
      </template>
      <template #in-playlist-trailing>
        <UIcon name="i-ph-bookmark-simple-fill" class="shrink-0 size-5" />
      </template>
    </UDropdownMenu>
  </div>
</template>
