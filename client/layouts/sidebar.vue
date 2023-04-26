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
                @other-selected="getOtherSelected"
                @toggle-layout="toggleLayout"
                :servers="filterServer"
                class="w-3/12"
                style="max-width: 280px"
              ></Sidebar>
              <Content
                v-if="element_selected !== 'show_create_mock'"
                :item="item_select"
                :item_link="item_link"
                :uri_params="uri_param"
                class="w-10/12"
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
import { useURIStore } from "~/stores/test";
//get all collection
const message = useMessage();
const dialog = useDialog();
const placement = "bottom-right";
// const { data: sidebarList } = await useFetch(
//   "http://127.0.0.1:8000/admin/collection.json",
//   {}
// );
const element_selected = useState("toggle_mock_collection");
const code_res = useState("code_response");
// const collections = ref(JSON.parse(JSON.stringify(sidebarList.value)));
const fakeData = [
  {
    id: "eea31c8d-5b83-44d8-bc79-e2a0e732f335",
    name: "alid",
    user_id: 1,
    requests: [
      {
        id: "ecdd233b-c8ad-427d-beca-59fc9c18c5c9",
        name: "/api/user",
        method: "GET",
        uri_component:
          "http://localhost:8000/eea31c8d-5b83-44d8-bc79-e2a0e732f335/api/user",
        query: "W3sia2V5IjoibGltaXQiLCJ2YWx1ZSI6IjEwIn1d",
        header: "W10=",
        body: "",
        collection_id: "eea31c8d-5b83-44d8-bc79-e2a0e732f335",
        responses: [
          {
            id: "c6dea9f7-b989-4c1b-af27-6ca4ac650b94",
            name: "/api/user",
            body: '{"user":2}',
            method: "GET",
            uri_component:
              "http://localhost:8000/eea31c8d-5b83-44d8-bc79-e2a0e732f335/api/user",
            request_id: "ecdd233b-c8ad-427d-beca-59fc9c18c5c9",
            created_at: "2023-04-22T02:34:49.553+07:00",
            updated_at: "2023-04-22T02:34:49.553+07:00",
          },
        ],
        created_at: "2023-04-22T02:34:49.526+07:00",
        updated_at: "2023-04-22T02:38:24.451+07:00",
      },
    ],
    is_server: 0,
    variable: "",
    created_at: "2023-04-22T02:34:49.166+07:00",
    updated_at: "2023-04-22T02:34:49.166+07:00",
  },
];
const collections = ref(fakeData);
const uriStore = useURIStore();
// const filterServer1 = collections.value.data.filter(
//   (server) => server.is_server == true
// );
// const filterServer = ref(filterServer1);
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
  uriStore.uri = uri_param.value;
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
    variable: "",
    is_server: true,
  };
  const request = {
    id: uuid_req,
    name: newapi.url,
    collection_id: collection.id,
    method: newapi.method,
    uri_component: "",
    responses: [],
    query: "",
    header: "",
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
const getOtherSelected = (item: Array) => {
  item_link.value = item;
  console.log("alo");
};
</script>

<style></style>
