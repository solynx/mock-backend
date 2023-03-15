<template>
  <NSpace vertical>
    <NInputGroup>
      <NSelect :options="selectOptions" v-model:value="Method" class="w-2/12" />

      <NInput :style="{ width: '70%' }" class="bg-neutral-200" />
      <NButton
        :style="{ width: '13%' }"
        :attr-type="submit"
        type="info"
        class="no-tailwind"
      >
        Send
      </NButton>
    </NInputGroup>
  </NSpace>
</template>

<script setup>
import { NSpace, NInputGroup, NSelect, NInput, NButton } from "naive-ui";
const Method = ref("GET");
const selectOptions = [
  { label: "GET", value: "GET", key: "key1" },
  { label: "POST", value: "POST", key: "key2", labelX: "labelX2" },
  { label: "PUT", value: "PUT", key: "key1", labelX: "labelX1" },
  { label: "PATCH", value: "PATCH", key: "key1", labelX: "labelX1" },
  { label: "DELETE", value: "DELETE", key: "key1", labelX: "labelX1" },
];
</script>

<style scoped>
.no-tailwind {
  background-color: #4098fc !important;
}
.no-tailwind:hover {
  opacity: 0.8;
}
</style>
