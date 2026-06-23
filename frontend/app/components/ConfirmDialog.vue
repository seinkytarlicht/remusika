<script setup lang="ts">
type PageProps = {
  title?: string;
  desc?: string;
  buttonYes?: string;
  buttonNo?: string;
  swap?: boolean;
};

const {
  title = "Are you sure?",
  desc,
  buttonNo = "No",
  buttonYes = "Yes",
  swap = false,
} = defineProps<PageProps>();

const emit = defineEmits<{ close: [boolean] }>();
</script>

<template>
  <UModal
    :close="{ onClick: () => emit('close', false) }"
    :title="title"
    :description="desc"
  >
    <template #footer>
      <div class="flex gap-2 items-end justify-end w-full min-h-[100px]">
        <UButton
          :label="buttonNo"
          @click="emit('close', false)"
          :color="swap ? 'neutral' : 'error'"
        />
        <UButton
          :label="buttonYes"
          @click="emit('close', true)"
          :color="swap ? 'error' : 'success'"
        />
      </div>
    </template>
  </UModal>
</template>
