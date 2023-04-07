<template>
  <NMessageProvider>
    <NDialogProvider>
      <NNotificationProvider :placement="placement">
        <NLoadingBarProvider>
          <NDialogProvider>
            <div class="wrapper flex">
              <Sidebar
                :collections="collections.data"
                @item-selected="getItemSelected"
                @toggle-layout="toggleLayout"
                :servers="filterServer"
                class="w-1/5"
              ></Sidebar>
              <Content
                v-if="element_selected !== 'show_create_mock'"
                :item="item_select"
                :item_link="item_link"
                :uri_params="uri_param"
                class="w-4/5"
              >
              </Content>
              <MockServer
                v-if="element_selected === 'show_create_mock'"
                @createServer="createMockServer"
                class="w-4/5"
              ></MockServer>
            </div>
          </NDialogProvider>
        </NLoadingBarProvider>
      </NNotificationProvider>
    </NDialogProvider>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import {
  NMessageProvider,
  NLoadingBarProvider,
  NCard,
  useMessage,
  useDialog,
  NNotificationProvider,
  NDialogProvider,
} from "naive-ui";
import { v4 as uuidv4 } from "uuid";
//get all collection
const message = useMessage();
const dialog = useDialog();
const placement = "bottom-right";
const { data: sidebarList } = await useFetch(
  "http://127.0.0.1:8000/admin/collection.json",
  {}
);
const element_selected = useState("toggle_mock_collection");
const code_res = useState("code_response");
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
  element_selected.value = "request";
  code_res.value = "";
  item_link.value = bread_cum;

  item_link.value.push({ id: item.id, name: item.name });
  console.log(item_link.value);
  item_select.value = item;

  uri_param.value = item.uri_component ?? "";
};
const modeCollection = ref(true);
const toggleLayout = () => {
  modeCollection.value = !modeCollection.value;
};
const createMockServer = async (newapi: object, collection_name: string) => {
  const uuid = uuidv4();
  const uuid_req = uuidv4();
  const uuid_res = uuidv4();
  const collection = {
    name: collection_name,
    id: uuid,
    requests: [],
    folders: [],

    is_server: true,
  };
  const request = {
    id: uuid_req,
    name: newapi.url,
    collection_id: collection.id,
    method: newapi.method,
    uri_component: "",
    responses: [],
  };

  const response = {
    id: uuid_res,
    name: request.name,
    request_id: request.id,
    method: request.method,
    uri_component: "",
    body: newapi.response_body,
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
    const api = {
      collection_id: collection.id,
      path: newapi.url,
      method: newapi.method,
      status_code: newapi.response_code,
      response_body: newapi.response_body,
    };
    if (await createMockApi(api)) {
      filterServer.value.push(collection);
      const url = "http://localhost:8000/" + collection.id;
      request.uri_component = url + request.name;
      response.uri_component = request.uri_component;
      const { data: status } = await useFetch(
        "http://localhost:8000/admin/request.json",
        {
          method: "POST",

          headers: { "Content-type": "application/json" },
          body: JSON.stringify(request),
        }
      );

      const result = JSON.parse(JSON.stringify(status.value));
      if (result.status) {
        const { data: status1 } = await useFetch(
          "http://localhost:8000/admin/response.json",
          {
            method: "POST",

            headers: { "Content-type": "application/json" },
            body: JSON.stringify(response),
          }
        );
        const result = JSON.parse(JSON.stringify(status1.value));
        if (result.status) {
          request.responses.push(response);
          collection.requests.push(request);
          collections.value.data.push(collection);
          return handleSuccess(url);
        }
      }
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
