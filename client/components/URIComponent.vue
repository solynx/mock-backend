<template>
  <NSpace vertical>
    <div style="display: none">
      {{ $test(props.uri.uri_component + $myPlugin.showUri()) }}
    </div>

    <NInputGroup>
      <NSelect
        :options="selectOptions"
        v-model:value="props.uri.method"
        class="w-2/12"
        size="large"
      />

      <NInput
        :style="{ width: '100%' }"
        class="bg-neutral-200 input-not-focus"
        v-model:value="a3"
        @keyup.enter="getRequest"
        @input="testChange"
        size="large"
      />
      <NButton
        :style="{ width: '13%' }"
        :attr-type="submit"
        type="info"
        class="no-tailwind"
        @click="getRequest"
        v-if="item_selected === 'request'"
        size="large"
      >
        Send
      </NButton>
    </NInputGroup>
  </NSpace>
</template>

<script setup>
import consolaGlobalInstance from "consola";
import { NSpace, NInputGroup, NSelect, NInput, NButton } from "naive-ui";

import { useFullUrlStore } from "~~/stores/uri_params";
const props = defineProps({
  uri: Object,
  filter_param: Array,
});

const urlFull = useFullUrlStore();
const { $test } = useNuxtApp();
const { $updateParams } = useNuxtApp();
const { $myPlugin } = useNuxtApp();
const query_actived = useState("query_active");
const abc = computed(() => props.uri.uri_component + $myPlugin.showUri());
const a1 = urlFull.getFullUrl(props.uri.uri_component + $myPlugin.showUri());
const query_1 = useState("data");
const a3 = useState("full_url_params");
const a = ref("");
a.value = $test(props.uri.uri_component + $myPlugin.showUri());
const item_selected = useState("item_select");
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
    link: a3.value,
    method: props.uri.method,
  });
};
// watch(a3, (value) => {
//   console.log(value);
//   const paramsString = value.split("?")[1];
//   const searchParams = new URLSearchParams(paramsString);

//   let query_temp = [];
//   let index = 0;
//   //Iterating the search parameters
//   searchParams.forEach((value, key) => {
//     query_temp.push({
//       key: query_actived.value[index].key,
//       param: key,
//       value: value,
//       description: "",
//     });
//     index++;
//   });
//   query_temp.sort((a, b) => a.key - b.key);
//   console.log(query_temp);
//   const updatedArr = query_1.value.map((item) => {
//     const index = query_temp.findIndex((obj) => obj.key === item.key);
//     if (index !== -1) {
//       return query_temp[index];
//     } else {
//       return item;
//     }
//   });

//   console.log("moi lan goi watch");
//   query_1.value = updatedArr;
//   console.log(query_1.value);
// });
const checkParam = () => {
  console.log(a3.value);
  return;
  const paramsString = a3.value.split("?")[1];

  const searchParams = new URLSearchParams(paramsString);

  let query_temp = [];
  let index = 0;
  //Iterating the search parameters
  searchParams.forEach((value, key) => {
    query_temp.push({
      key: query_actived.value[index].key,
      param: key,
      value: value,
      description: "",
    });
    index++;
  });
  query_temp.sort((a, b) => a.key - b.key);

  // for (let i = 0; i < query_temp.length; i++) {
  //   const index = query_1.value.findIndex(
  //     (item) => item.key === query_temp[i].key
  //   );
  //   if (index !== -1) {
  //     query_1.value[index] = query_temp[i];
  //   }
  // }

  // query_1.value = updatedArr;
  console.log(query_1.value);
};
const testChange = () => {
  const paramsString = a3.value.split("?")[1];

  const searchParams = new URLSearchParams(paramsString);

  let query_temp = [];
  let index = 0;
  //Iterating the search parameters
  if (query_actived.value.length) {
    searchParams.forEach((value, key) => {
      query_temp.push({
        key: query_actived.value[index].key,
        param: key,
        value: value,
        description: "",
      });
      index++;
    });
    query_temp.sort((a, b) => a.key - b.key);
    for (let i = 0; i < query_temp.length; i++) {
      const index = query_1.value.findIndex(
        (item) => item.key === query_temp[i].key
      );
      const index2 = query_actived.value.findIndex(
        (item) => item.key === query_temp[i].key
      );
      if (index !== -1) {
        query_1.value[index] = query_temp[i];
      }
      if (index2 !== -1) {
        query_actived.value[index] = query_temp[i];
      }
    }
  }
  //else {
  //   let index = 0;
  //   searchParams.forEach((value, key) => {
  //     query_temp.push({
  //       key: index + 1,
  //       param: key,
  //       value: value,
  //       description: "",
  //     });
  //     index++;
  //   });
  //   for (let i = 0; i < query_temp.length; i++) {
  //     query_1.value[i] = query_temp[i];

  //     query_actived.value[i] = query_temp[i];
  //   }
  //   query_1.value.push({key:query_temp.length,param:"",value:"",description:""});
  // }
};
</script>

<style scoped>
.no-tailwind {
  background-color: #4098fc !important;
}
.no-tailwind:hover {
  opacity: 0.8;
}
.input-not-focus {
  background-color: #f2f2f2;
}
</style>
