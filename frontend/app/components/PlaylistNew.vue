<script setup lang="ts">
import type { Form, FormSubmitEvent } from "@nuxt/ui";
import * as z from "zod";

const { $api } = useNuxtApp();
const toast = useToast();
const playlistStore = usePlaylistStore();

const formRef = ref<Form<Schema>>();

const schema = z.object({
  name: z
    .string("Please put playlist name")
    .trim()
    .min(1, "Please put playlist name"),
});

type Schema = z.output<typeof schema>;

const playlistModel = reactive<Partial<Schema>>({
  name: undefined,
});

async function savePlaylist(
  event: FormSubmitEvent<Schema>,
  callbakc: () => void,
) {
  try {
    await $api("/playlist/create", {
      method: "POST",
      body: event.data,
    });

    toast.add({
      title: "Playlist created",
    });
    playlistStore.fetch();
    playlistModel.name = "";
    callbakc();
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
</script>

<template>
  <UModal
    title="Create new playlist"
    description="You can add music after create it"
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
      <UForm
        :schema="schema"
        :state="playlistModel"
        class="space-y-4"
        @submit="(e) => savePlaylist(e, close)"
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
            <UButton label="Cancel" color="error" @click="close" />
            <UButton label="Create" color="success" type="submit" />
          </div>
        </div>
      </UForm>
    </template>
  </UModal>
</template>
