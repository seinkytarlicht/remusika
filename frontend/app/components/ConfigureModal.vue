<script setup lang="ts">
import type { TabsItem } from "@nuxt/ui";
import VerticalInputGroup from "./VerticalInputGroup.vue";

const tabs = ref<TabsItem[]>([
  {
    label: "Setting",
    icon: "i-lucide-settings",
    slot: "setting",
  },
  {
    label: "Help",
    icon: "i-lucide-circle-question-mark",
    slot: "help",
  },
  {
    label: "About",
    icon: "i-lucide-info",
    slot: "about",
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
    <UButton icon="i-lucide-settings-2" color="neutral" variant="outline" />

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
              <VerticalInputGroup text="Play/Pause Video" class="py-2">
                <UKbd value="k" size="lg" />
              </VerticalInputGroup>
              <VerticalInputGroup text="Next Video" class="py-2">
                <div class="flex">
                  <UKbd size="lg">shift</UKbd>
                  <UKbd value="n" size="lg" />
                </div>
              </VerticalInputGroup>
              <VerticalInputGroup text="Prev Video" class="py-2">
                <div class="flex">
                  <UKbd size="lg">shift</UKbd>
                  <UKbd size="lg">alt</UKbd>
                  <UKbd value="n" size="lg" />
                </div>
              </VerticalInputGroup>
              <VerticalInputGroup text="Loop Video" class="py-2">
                <div class="flex">
                  <UKbd value="l" size="lg" />
                </div>
              </VerticalInputGroup>
            </main>
          </template>

          <template #about>
            <img src="/rem.png" class="size-32 mb-2" />
            <Logo />
            <p class="text-muted">Build version 0.1</p>
          </template>
        </UTabs>
      </div>
    </template>
  </UModal>
</template>
