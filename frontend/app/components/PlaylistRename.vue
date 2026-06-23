<script setup lang="ts">
import type { Playlist } from "~/types/playlist";

const { $api } = useNuxtApp();
const toast = useToast();
const playlistStore = usePlaylistStore();

type PageProps = {
  playlist: Playlist;
};

const { playlist } = defineProps<PageProps>();

const playlistModel = reactive<{ id: number; name: string }>({
  id: playlist.id,
  name: playlist.name,
});

async function savePlaylist() {
  try {
    await $api("/playlist/update", {
      method: "PUT",
      body: playlistModel,
    });

    toast.add({
      title: "Playlist updated",
    });
    playlistStore.fetch();
    playlistModel.name = "";
    emit("close", true);
  } catch (e) {
    console.error(e);
    toast.add({
      title: "Failed to update playlist",
      description: "Error: " + String(e),
      color: "error",
    });
  }
}

const emit = defineEmits<{ close: [boolean] }>();
</script>

<template>
  <UModal
    title="Create new playlist"
    description="You can add music after create it"
    :dismissible="false"
  >
    <template #body>
      <div class="flex flex-col gap-6 min-h-50">
        <UFormField label="Playlist Name" required>
          <UInput
            type="text"
            v-model="playlistModel.name"
            placeholder="E.g., Hardbass, Nightcore, Gym Music, 67, etc"
            class="w-full"
          />
        </UFormField>

        <div class="flex gap-2 justify-end w-full mt-auto">
          <UButton label="Cancel" color="error" @click="emit('close', false)" />
          <UButton label="Save" color="success" @click="savePlaylist" />
        </div>
      </div>
    </template>
  </UModal>
</template>
