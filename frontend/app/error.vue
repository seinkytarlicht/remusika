<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps<{ error: NuxtError }>();

const router = useRouter();

onMounted(() => {
  router.back();
});
</script>

<template></template>
