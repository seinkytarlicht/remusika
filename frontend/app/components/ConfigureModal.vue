<script setup lang="ts">
import type { TabsItem } from "@nuxt/ui";
import VerticalInputGroup from "./VerticalInputGroup.vue";
const config = useRuntimeConfig();

const tabs = ref<TabsItem[]>([
  {
    label: "Setting",
    icon: "i-ph-gear-six",
    slot: "setting",
  },
  {
    label: "Help",
    icon: "i-ph-question",
    slot: "help",
  },
  {
    label: "About",
    icon: "i-ph-info",
    slot: "about",
  },
]);

const shortcutLists = ref([
  {
    title: "Play/Pause Video",
    kbd: ["K", "Space"],
  },
  {
    title: "Skip Forward",
    kbd: ["arrow right"],
  },
  {
    title: "Rewind",
    kbd: ["arrow left"],
  },
  {
    title: "Next Video",
    kbd: ["shift + n"],
  },
  {
    title: "Prev Video",
    kbd: ["shift + alt + n"],
  },
  {
    title: "Loop Video",
    kbd: ["l"],
  },
]);
</script>

<template>
  <UModal
    title="Configure"
    :ui="{
      content: 'max-w-3xl',
    }"
  >
    <UButton icon="i-ph-gear-six" color="neutral" variant="outline" />

    <template #body>
      <div class="h-125 max-h-125">
        <UTabs
          orientation="vertical"
          :items="tabs"
          class="w-full"
          variant="link"
          :ui="{
            root: 'items-start gap-8',
            content: 'mt-2',
          }"
        >
          <template #setting>
            <main class="h-full flex flex-col divide-y divide-accented">
              <VerticalInputGroup
                text="Use Dark Mode"
                description="You're monsters if you use light mode..."
                class="py-2"
              >
                <UColorModeSwitch size="xl" />
              </VerticalInputGroup>
            </main>
          </template>

          <template #help>
            <h3 class="text-xl font-bold">Shortcuts</h3>
            <main class="h-full flex flex-col divide-y divide-accented">
              <VerticalInputGroup
                :text="s.title"
                class="py-2"
                v-for="s in shortcutLists"
              >
                <div class="flex gap-2">
                  <UKbd size="lg" v-for="kbd in s.kbd">{{ kbd }}</UKbd>
                </div>
              </VerticalInputGroup>
            </main>
          </template>

          <template #about>
            <img src="/rem.png" class="size-32 mb-2" />
            <Logo />
            <p class="text-muted">
              Build version {{ config.public.versionBuild }}
            </p>
          </template>
        </UTabs>
      </div>
    </template>
  </UModal>
</template>
