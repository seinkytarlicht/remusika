<script setup lang="ts">
import type { Form, FormSubmitEvent } from "@nuxt/ui";
import * as z from "zod";
import type { Playlist } from "~/types/playlist";

const { $api } = useNuxtApp();
const toast = useToast();
const playlistStore = usePlaylistStore();

type PageProps = {
  playlist: Playlist;
};

const { playlist } = defineProps<PageProps>();

const formRef = ref<Form<Schema>>();

const schema = z.object({
  id: z.number().min(1),
  name: z
    .string("Please put playlist name")
    .trim()
    .min(1, "Please put playlist name"),
});

type Schema = z.output<typeof schema>;

const playlistModel = reactive<Partial<Schema>>({
  id: playlist.id,
  name: playlist.name,
});

async function savePlaylist(event: FormSubmitEvent<Schema>) {
  try {
    await $api("/playlist/update", {
      method: "PUT",
      body: event.data,
    });

    toast.add({
      title: "Playlist updated",
    });
    playlistStore.fetch();
    emit("close", true);
  } catch (e: any) {
    if (e.statusCode != 400) {
      toast.add({
        title: "Failed to add playlist",
        description: "Error: See console for more info",
        color: "error",
      });
      console.error(e);
      return;
    }

    formRef.value?.setErrors([
      {
        name: "name",
        message: e.data.messages ?? "Error",
      },
    ]);
  }
}

const emit = defineEmits<{ close: [boolean] }>();
</script>

<template>
  <UModal
    :close="{ onClick: () => emit('close', false) }"
    title="Rename playlist"
  >
    <template #body>
      <UForm
        :schema="schema"
        :state="playlistModel"
        class="space-y-4"
        @submit="savePlaylist"
        ref="formRef"
      >
        <div class="flex flex-col gap-6 min-h-50">
          <UFormField label="Playlist Name" name="name">
            <UInput
              type="text"
              v-model="playlistModel.name"
              placeholder="E.g., Lofi, Nightcore, Gym Music, 3 AM, 67, etc"
              class="w-full"
            />
          </UFormField>

          <div class="flex gap-2 justify-end w-full mt-auto">
            <UButton
              label="Cancel"
              color="error"
              @click="emit('close', false)"
            />
            <UButton label="Rename" color="success" type="submit" />
          </div>
        </div>
      </UForm>
    </template>
  </UModal>
</template>
