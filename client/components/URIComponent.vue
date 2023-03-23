<template>
  <NSpace vertical>
    {{ $myPlugin() }}
    <input v-model="test12.value" />
    {{ test12.value }}
    <NInputGroup>
      <NSelect
        :options="selectOptions"
        v-model:value="props.uri.method"
        class="w-2/12"
      />
      <NInput
        :style="{ width: '70%' }"
        class="bg-neutral-200"
        v-model:value="props.uri.uri_component"
        @keyup.enter="getRequest"
      />
      <NButton
        :style="{ width: '13%' }"
        :attr-type="submit"
        type="info"
        class="no-tailwind"
        @click="getRequest"
      >
        Send
      </NButton>
    </NInputGroup>
  </NSpace>
</template>

<script setup>
import { NSpace, NInputGroup, NSelect, NInput, NButton } from "naive-ui";
const props = defineProps({
  uri: Object,
  filter_param: Array,
});
const { $myPlugin } = useNuxtApp();

const same_counter = useUriQP();
const test12 = ref($myPlugin());
console.log(test12.value);
const qp = same_counter.value;
const emit = defineEmits(["change_link"]);
const Method = ref("GET");
const upperCaseMethod = computed(() =>
  props.uri.method ? props.uri.method.toUpperCase() : null
);
const selectOptions = [
  { label: "GET", value: "GET", key: "key1" },
  { label: "POST", value: "POST", key: "key2", labelX: "labelX2" },
  { label: "PUT", value: "PUT", key: "key1", labelX: "labelX1" },
  { label: "PATCH", value: "PATCH", key: "key1", labelX: "labelX1" },
  { label: "DELETE", value: "DELETE", key: "key1", labelX: "labelX1" },
];
const getRequest = () => {
  emit("change_link", {
    link: props.uri.uri_component,
    method: props.uri.method,
  });
};
</script>

<style scoped>
.no-tailwind {
  background-color: #4098fc !important;
}
.no-tailwind:hover {
  opacity: 0.8;
}
</style>
