<script setup lang="ts">
import type { Playlist } from "~/types/playlist";
import ConfirmDialog from "./ConfirmDialog.vue";
import PlaylistRename from "./PlaylistRename.vue";

type PageProps = {
  playlist: Playlist;
  dropdown?: boolean;
};

const { playlist: pl, dropdown = true } = defineProps<PageProps>();

const { $api } = useNuxtApp();
const route = useRoute();
const isUseDrawer = useLocalStorage("remusika_use_drawer", false);
const overlay = useOverlay();
const toast = useToast();
const playlistStore = usePlaylistStore();

const modalConfirm = overlay.create(ConfirmDialog);
const modalRename = overlay.create(PlaylistRename);
const dropdownModel = ref(false);

async function deletePlaylist() {
  const confirmDialog = modalConfirm.open({
    title: `Are you sure want to delete "${pl.name}" playlist?`,
    desc: "You can't undo this action.",
    buttonNo: "No! I want to keep it!",
    buttonYes: "Yes! delete it!",
    swap: true,
  });

  const isConfirm = await confirmDialog.result;

  if (isConfirm) {
    try {
      await $api(`/playlist/delete/${pl.id}`, {
        method: "DELETE",
      });

      toast.add({
        title: "Playlist deleted",
      });
      playlistStore.fetch();
    } catch (e) {
      console.error(e);
      toast.add({
        title: "Failed to delete playlist",
        description: String(e),
      });
    }
  }
}

async function renamePlaylist() {
  modalRename.open({
    playlist: pl,
  });
}
</script>

<template>
  <div
    class="flex items-center justify-between gap-5 py-2 px-3 rounded-sm relative min-w-0"
    v-on:contextmenu="
      (e) => {
        e.preventDefault();
        dropdownModel = true;
      }
    "
    :class="
      !isUseDrawer
        ? Number(route.query['showlist']) === pl.id
          ? 'bg-primary/20 text-primary '
          : 'hover:bg-primary/5'
        : Number(route.query['playlist']) === pl.id
          ? 'bg-primary/20 text-primary '
          : 'hover:bg-primary/5'
    "
  >
    <NuxtLink
      :to="{
        query: {
          ...route.query,
          showlist: pl.id,
        },
      }"
      @click="isUseDrawer = false"
      class="absolute bottom-0 left-0 top-0 right-0"
    >
    </NuxtLink>

    <p class="truncate">
      {{ pl.name }}
    </p>

    <UDropdownMenu
      v-if="dropdown"
      v-model:open="dropdownModel"
      :modal="false"
      :items="[
        {
          label: 'Rename',
          icon: 'i-ph-pencil-simple-line-fill',
          onSelect: renamePlaylist,
        },
        {
          label: 'Delete',
          icon: 'i-ph-trash-simple-fill',
          color: 'error',
          onSelect: deletePlaylist,
        },
      ]"
      :content="{
        align: 'start',
      }"
      :ui="{
        content: 'min-w-48 max-h-54',
      }"
    >
      <UButton
        class="relative p-0"
        icon="i-ph-dots-three-outline-vertical-fill"
        color="neutral"
        variant="link"
        size="sm"
      />
    </UDropdownMenu>
  </div>
</template>
