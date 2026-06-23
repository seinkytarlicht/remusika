<script setup lang="ts">
import type { Playlist } from "~/types/playlist";
import ConfirmDialog from "./ConfirmDialog.vue";

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

const modal = overlay.create(ConfirmDialog);
const dropdownModel = ref(false);

async function deletePlaylist() {
  const confirmDialog = modal.open({
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
</script>

<template>
  <div
    class="flex items-center justify-between py-2 px-3 rounded-sm relative"
    v-on:contextmenu="
      (e) => {
        e.preventDefault();
        dropdownModel = true;
      }
    "
    :class="
      route.query.showlist === pl.name
        ? 'bg-primary/20 text-primary '
        : 'hover:bg-primary/5'
    "
  >
    <NuxtLink
      :to="{
        query: {
          ...route.query,
          showlist: pl.name,
        },
      }"
      @click="isUseDrawer = false"
      class="absolute bottom-0 left-0 top-0 right-0"
    >
    </NuxtLink>

    <p>
      {{ pl.name }}
    </p>

    <UDropdownMenu
      v-if="dropdown"
      v-model:open="dropdownModel"
      :modal="false"
      :items="[
        {
          label: 'Delete Playlist',
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
