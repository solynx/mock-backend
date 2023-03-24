<template>
  <NMessageProvider>
    <NLoadingBarProvider>
      <div class="wrapper flex">
        <Sidebar
          :collections="collections.data"
          @item-selected="getItemSelected"
        ></Sidebar>
        <Content
          :item="item_select"
          :item_link="item_link"
          :uri_params="uri_param"
        ></Content>
      </div>
    </NLoadingBarProvider>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import { NMessageProvider, NLoadingBarProvider } from "naive-ui";
//get all collection
const { data: sidebarList } = await useFetch(
  "http://127.0.0.1:8000/admin/collection.json",
  {}
);

const collections = ref(JSON.parse(JSON.stringify(sidebarList.value)));

const item_select = ref({});
const item_link = ref([]);
const uri_param = ref("");
const getItemSelected = (bread_cum: Array, item: object) => {
  item_link.value = bread_cum;
  item_link.value.push({ id: item.id, name: item.name });

  item_select.value = item;
  uri_param.value = item.uri_component;
  console.log(uri_param);
};
</script>

<style></style>
