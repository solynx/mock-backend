<template>
  <NMessageProvider>
    <div class="wrapper flex">
      <Sidebar
        :collections="collections"
        @item-selected="getItemSelected"
      ></Sidebar>
      <Content
        :item="item_select"
        :item_link="item_link"
        :uri_params="uri_param"
      ></Content>
    </div>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import { NMessageProvider } from "naive-ui";
var collections1 = [
  {
    name: "New Collection",
    id: "col_abcxyz111",
    folders: [
      {
        name: "folder1",
        id: "aabbcc12",
        requests: [],
      },
      {
        name: "folder2",
        id: "fd_aabbccaz",
        requests: [
          {
            name: "Request1",
            method: "Post",
            id: "req_ndnd",
            uri_component: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                id: "res_aksndasndf",
              },
              {
                name: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                id: "res_aksdkas2",
              },
            ],
          },
          {
            name: "Request123",
            method: "Post",
            id: "req_ndnd",
            uri_component: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                id: "res_aksndasnd12",
              },
              {
                name: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                id: "res_aksdkas1",
              },
            ],
          },
        ],
      },
    ],
    requests: [
      {
        name: "Request A",
        id: "clt_req_1",
        method: "Patch",
        uri_component: "facebook.com",
        responses: [
          {
            name: "Response 1",
            id: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
          {
            name: "Response 1",
            id: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
        ],
      },
    ],
  },
  {
    name: "New Collection",
    id: "col_abcxyz45",
    folders: [
      {
        name: "folder1",
        id: "aabbcc12354",
      },
      {
        name: "folder2",
        id: "fd_aabbccaz234",
        requests: [
          {
            name: "Request1",
            method: "Post",
            id: "req_ndnd234",
            uri_component: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                id: "res_aksndasndf1",
              },
              {
                name: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                id: "res_aksdkas21",
              },
            ],
          },
          {
            name: "Request123",
            method: "Post",
            id: "req_ndnd12",
            uri_component: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                id: "res_aksndasnd123",
              },
              {
                name: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                id: "res_aksdkas12",
              },
            ],
          },
        ],
      },
    ],
  },
  {
    name: "New Collection2",
    id: "abcxyz1",
  },
];
const collections = ref(collections1);
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
