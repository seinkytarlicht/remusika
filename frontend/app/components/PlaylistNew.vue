<script setup lang="ts">
const { $api } = useNuxtApp();
const toast = useToast();
const playlistStore = usePlaylistStore();

const playlistModel = reactive<{ name: string }>({
  name: "",
});

async function savePlaylist(callback: () => void) {
  try {
    await $api("/playlist/create", {
      method: "POST",
      body: playlistModel,
    });

    toast.add({
      title: "Playlist created",
    });
    playlistStore.fetch();
    playlistModel.name = "";
    callback();
  } catch (e) {
    console.error(e);
    toast.add({
      title: "Failed to add playlist",
      description: "Error: " + String(e),
      color: "error",
    });
  }
}
</script>

<template>
  <UModal
    title="Create new playlist"
    description="You can add music after create it"
    :dismissible="false"
  >
    <UTooltip
      text="Add playlist"
      :delay-duration="0"
      :content="{
        side: 'top',
      }"
    >
      <UButton icon="i-ph-plus-bold" variant="soft" />
    </UTooltip>

    <template #body="{ close }">
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
          <UButton label="Cancel" color="error" @click="close" />
          <UButton
            label="Create"
            color="success"
            @click="() => savePlaylist(close)"
          />
        </div>
      </div>
    </template>
  </UModal>
</template>
