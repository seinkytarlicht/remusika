<script setup lang="ts">
type PageProps = {
  text: string;
  description?: string;
};

const props = defineProps<PageProps>();
</script>

<template>
  <div class="flex justify-between items-center">
    <div>
      <h4 class="">{{ props.text }}</h4>
      <p class="text-muted text-sm">{{ props.description }}</p>
    </div>
    <slot />
  </div>
</template>
