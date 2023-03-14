<template>
  <div class="flex">
    <n-select
      :options="selectOptions"
      v-model:value="HttpType"
      class="w-2/12"
    />
    <n-select
      :options="selectDataTypeInRawOfBody"
      v-model:value="DataType"
      class="w-1/12"
    />
  </div>
</template>

<script lang="ts" setup>
import { defineComponent, ref } from "vue";
import { SelectOption, SelectGroupOption, NSelect } from "naive-ui";
const HttpType = ref("None");
const DataType = ref("text");
const selectDataTypeInRawOfBody = [
  { label: "Text", value: "text", key: "key11" },
  { label: "JavaScript", value: "js", key: "key12" },
  { label: "JSON", value: "json", key: "key13" },
  { label: "HTML", value: "html", key: "key14" },
  { label: "XML", value: "xml", key: "key15" },
];

const selectOptions = [
  { label: "none", value: "None", key: "key11" },
  { label: "form-data", value: "POST", key: "form-data" },
  {
    label: "x-www-form-urlencoded",
    value: "urlencoded",
    key: "key12",
  },
  {
    label: "raw",
    value: "raw",
    key: "key13",
  },
  { label: "binary", value: "binary", key: "key14" },
  { label: "GraphQL", value: "graphql", key: "key15" },
];
</script>
