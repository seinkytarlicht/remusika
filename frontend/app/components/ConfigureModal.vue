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
              >
                <UColorModeSwitch size="xl" />
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
