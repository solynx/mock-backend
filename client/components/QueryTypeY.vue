<template>
  <div class="flex">
    <n-select
      :options="selectOptions"
      v-model:value="HttpType"
      class="w-5/12"
    />
    <n-select
      v-if="HttpType === 'raw'"
      :options="selectDataTypeInRawOfBody"
      v-model:value="DataType"
      class="w-3/12"
    />
  </div>
  <BodyLayer :layer="HttpType" />
</template>

<script lang="ts" setup>
import { defineComponent, ref } from "vue";
import { SelectOption, SelectGroupOption, NSelect } from "naive-ui";
const HttpType = ref("none");
const DataType = ref("json");
const selectDataTypeInRawOfBody = [
  { label: "Text", value: "text", key: "key11" },
  { label: "JavaScript", value: "js", key: "key12" },
  { label: "JSON", value: "json", key: "key13" },
  { label: "HTML", value: "html", key: "key14" },
  { label: "XML", value: "xml", key: "key15" },
];

const selectOptions = [
  { label: "none", value: "none", key: "key11" },
  { label: "form-data", value: "form-data", key: "form-data" },
  {
    label: "x-www-form-urlencoded",
    value: "x-www-form-urlencoded",
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
