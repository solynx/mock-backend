<template>
  <NMessageProvider>
    <NLoadingBarProvider>
      <NDialogProvider>
        <div class="wrapper flex">
          <Sidebar
            :collections="collections.data"
            @item-selected="getItemSelected"
            @toggle-layout="toggleLayout"
            :servers="filterServer"
          ></Sidebar>
          <Content
            v-if="modeCollection"
            :item="item_select"
            :item_link="item_link"
            :uri_params="uri_param"
          >
          </Content>
          <MockServer v-else @createServer="createMockServer"></MockServer>
        </div>
      </NDialogProvider>
    </NLoadingBarProvider>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import {
  NMessageProvider,
  NLoadingBarProvider,
  NCard,
  useMessage,
  useDialog,
} from "naive-ui";
import { v4 as uuidv4 } from "uuid";
//get all collection
const message = useMessage();
const dialog = useDialog();
const { data: sidebarList } = await useFetch(
  "http://127.0.0.1:8000/admin/collection.json",
  {}
);

const collections = ref(JSON.parse(JSON.stringify(sidebarList.value)));
console.log(collections.value.data);
const filterServer1 = collections.value.data.filter(
  (server) => server.is_server == true
);

const handleSuccess = (url: string) => {
  dialog.success({
    title: "Create Mock Server Success!",
    content:
      "Send a request to the following mock server URL, followed by the request path: " +
      url,
    positiveText: "Copy URL",
    negativeText: "Cancel",
    onPositiveClick: () => {
      navigator.clipboard.writeText(url);
      message.success("URL Copied!");
    },
  });
};
const filterServer = ref(filterServer1);
const item_select = ref({});
const item_link = ref<object[]>([]);
const uri_param = ref("");
const getItemSelected = (bread_cum: Array, item: object) => {
  item_link.value = bread_cum;

  item_link.value.push({ id: item.id, name: item.name });
  console.log(item_link.value);
  item_select.value = item;

  uri_param.value = item.uri_component;
};
const modeCollection = ref(true);
const toggleLayout = () => {
  modeCollection.value = !modeCollection.value;
};
const createMockServer = async (newapi: object, collection_name: string) => {
  const uuid = uuidv4();
  const collection = {
    name: collection_name,
    id: uuid,
    requests: [],
    folders: [],

    is_server: true,
  };

  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/mock-api/new",
    {
      method: "POST",
      mode: "cors",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(collection),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    collections.value.data.push(collection);
    const api = {
      collection_id: collection.id,
      path: newapi.url,
      method: newapi.method,
      status_code: newapi.response_code,
      response_body: newapi.response_body,
    };
    if (await createMockApi(api)) {
      filterServer.value.push(collection);
      const url = "localhost:8000/" + collection.id;

      return handleSuccess(url);
    }
    return message.error("Failed!");
  }

  return message.error("Failed!");
};
const createMockApi = async (api: object) => {
  const method = api.method === "GET" ? "POST" : api.method;
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/mock-api/update",
    {
      method: method,
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(api),
    }
  );

  const status1 = JSON.parse(JSON.stringify(status.value));

  return status1.status;
};
</script>

<style></style>
