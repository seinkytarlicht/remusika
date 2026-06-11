<script setup lang="ts">
type PageProps = {
  text: string;
  description?: string;
};

const props = defineProps<PageProps>();
</script>

<template>
  <div class="flex justify-between items-center py-4">
    <div>
      <h4 class="font-semibold">{{ props.text }}</h4>
      <p class="text-muted text-sm">{{ props.description }}</p>
    </div>
    <slot />
  </div>
</template>
