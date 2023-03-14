<template>
  <NSpace vertical>
    <NInputGroup>
      <NSelect :options="selectOptions" class="w-2/12" />

      <NInput :style="{ width: '70%' }" />
      <NButton :style="{ width: '13%' }" type="info"> Send </NButton>
      <n-button type="primary"> Primary </n-button>
    </NInputGroup>
  </NSpace>
</template>

<script setup>
import { NSpace, NInputGroup, NSelect, NInput, NButton } from "naive-ui";

const selectOptions = [
  { label: "GET", value: "GET", key: "key1" },
  { label: "POST", value: "POST", key: "key2", labelX: "labelX2" },
  { label: "PUT", value: "PUT", key: "key1", labelX: "labelX1" },
  { label: "PATCH", value: "PATCH", key: "key1", labelX: "labelX1" },
  { label: "DELETE", value: "DELETE", key: "key1", labelX: "labelX1" },
];
</script>

<style></style>
