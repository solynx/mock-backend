<template>
  <div class="w-full" style="padding: 3px 10px">
    <!-- Phần tag multiple bao gồm có nhiều file được mở cùng lúc tại  thanh này-->
    <div class="w-no-x">
      <div class="file-multiple">
        <n-dynamic-tags v-model:value="tags" :max="5" />
      </div>
      <n-card size="small">
        <div class="file-detail flex justify-between">
          <div class="file-link">
            <FileLink :link="item_link"></FileLink>
          </div>
          <FileOption
            @update_response="updateResponseBody"
            @update_api="updateMockApi"
          ></FileOption>
        </div>
      </n-card>

      <URIComponent
        :uri="item"
        @change_link="getRequest"
        :filter_param="filteredQueryParam.value"
        v-if="item_selected === 'response' || item_selected === 'request'"
      ></URIComponent>
    </div>
    <!-- collection-->
    <div>
      <CollectionOptionTable
        v-if="item_selected === 'collection'"
      ></CollectionOptionTable>
    </div>
    <div
      class="group-func flex pt-4"
      v-if="item_selected === 'response' || item_selected === 'request'"
    >
      <n-card style="width: 50%; min-width: 40%">
        <div class="group-fun-query mt-3">
          <div class="request-option">
            <QueryDataType></QueryDataType>
          </div>
        </div>
      </n-card>

      <n-config-provider
        style="width: 50%; max-width: 50%"
        :hljs="hljs"
        v-if="item_selected === 'response' || item_selected === 'request'"
      >
        <n-loading-bar-provider>
          <NSpin :show="!disabledRef" size="small">
            <ResponseData v-if="item_selected === 'request'"></ResponseData>
            <MockData
              v-if="item_selected === 'response'"
              @update_response="updateResponseBody"
              @update_api="updateMockApi"
            ></MockData>
            <template #description> Just moment... </template>
          </NSpin>
        </n-loading-bar-provider>
      </n-config-provider>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  SaveOutline as SaveIcon,
  ChevronDownOutline as ChevronDownIcon,
  EllipsisHorizontalOutline as MoreIcon,
  PencilOutline as PencilIcon,
  ChatboxEllipsesOutline as MomentModeIcon,
} from "@vicons/ionicons5";
import {
  NSpace,
  NInput,
  NInputGroup,
  NButton,
  NDatePicker,
  NSelect,
  NCalendar,
  NTimePicker,
  NCard,
  NTabPane,
  NTabs,
  NTab,
  NTable,
  NCheckbox,
  NCheckboxGroup,
  NCode,
  NConfigProvider,
  NIcon,
  NDynamicTags,
  NTransfer,
  NScrollbar,
  NRadio,
  useLoadingBar,
  NSpin,
  useMessage,
} from "naive-ui";
import hljs from "highlight.js/lib/core";
import json from "highlight.js/lib/languages/json";
hljs.registerLanguage("json", json);

const props = defineProps({
  item: Object,
  item_link: Array,
});
const message = useMessage();
const loadingBar = useLoadingBar();
const disabledRef = ref(true);
const cascaderOptions = [
  {
    label: "option-1",
    value: "option-1",
    children: [
      {
        label: "option-1-1",
        value: "option-1-1",
      },
    ],
  },
];
const item_selected = useState("item_select");
const code = useState("code_response", () => "");
const header_actived = useState("header_actived");
const json_raw = useState("json_raw");
const var_actived = useState("var_actived", () => {});
const getRequest = async (request_detail: object) => {
  props.item.uri_component = request_detail.link;
  const collection_var = useState("var_collection");

  const params = useState("data");
  const header_req = useState("header_req");
  let new_params = [];
  let new_headers = [];
  const getUrl = props.item.uri_component.split("?")[0];
  props.item.uri_component = getUrl;

  params.value.forEach((item) => {
    if (item.param === "" && item.value === "") return;
    let param = {
      key: item.param,
      value: item.value,
    };
    new_params.push(param);
  });
  header_req.value.forEach((item) => {
    if (item.param === "" && item.value === "") return;
    let param = {
      key: item.param,
      value: item.value,
    };
    new_headers.push(param);
  });
  const jsonString = JSON.stringify(new_params);
  const base64String = btoa(jsonString);
  const jsonHeaderString = JSON.stringify(new_headers);
  const base64HeaderString = btoa(jsonHeaderString);
  loadingBar.start();
  disabledRef.value = false;
  let result;
  let temp_uri = props.item.uri_component;

  var regex = /{{([^}]+)}}/g;
  var match;
  while ((match = regex.exec(request_detail.link))) {
    var paramName = match[1];
    var paramValue = null;
    for (var i = 0; i < collection_var.value.length; i++) {
      if (collection_var.value[i].param === paramName) {
        paramValue = collection_var.value[i].value;
        break;
      }
    }
    if (paramValue) {
      request_detail.link = request_detail.link.replace(match[0], paramValue);
    }
  }

  const { data: status } = await useFetch(request_detail.link, {
    method: request_detail.method,
    headers: header_actived.value ?? { "Content-type": "application/json" },
    body: request_detail.method !== "GET" ? props.item.body : null,
  });
  let old_header = props.item.header;
  let old_query = props.item.query;
  props.item.header = base64HeaderString;
  props.item.query = base64String;

  const { data: result1 } = await useFetch(
    "http://localhost:8000/admin/request.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(props.item),
    }
  );

  result = JSON.parse(JSON.stringify(status.value));
  let result_save_req = JSON.parse(JSON.stringify(result1.value));

  if (!result_save_req.status) {
    props.item.uri_component = temp_uri;
    props.item.header = old_header;
    props.item.query = old_query;
  }

  loadingBar.finish();
  disabledRef.value = true;
  let bad_request = {
    error: 404,
    message: "The API Not Found!",
  };
  result = result !== null ? result : bad_request;
  code.value = JSON.stringify(result, null, 1);

  return;
};
const uri_param = ref("");
const filteredQueryParam = ref([]);
const toggleQueryParam = (activeQuery: Array, queryParamList: Array) => {
  // activeQuery.sort((a, b) => a - b);
  // const filteredValues = queryParamList.filter((value, index) => {
  //   return activeQuery.includes(index);
  // });
  // filteredQueryParam.value = filteredValues;
  // return updateUri(filteredQueryParam.value);

  filteredQueryParam.value = queryParamList;
};
const updateUri = (params: Array) => {
  filteredQueryParam.value = params;
  const query_params = useState("");
};
const tags = useState("tags_click");
const checkedValueRef = ref<string | null>(null);
const checkedValue = checkedValueRef;
const response_saving = useState("response_saving");
const collection_parent = useState("belong_collection");
const updateResponseBody = async (response: object) => {
  response_saving.value = true;
  // console.log(response_saving.status);
  // response_saving.opacity = true;
  loadingBar.start();
  disabledRef.value = false;
  let result;
  const { data: status } = await useFetch(
    "http://localhost:8000/admin/response.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(response),
    }
  );
  result = JSON.parse(JSON.stringify(status.value));
  loadingBar.finish();
  disabledRef.value = true;
  response_saving.value = false;
  if (result.status) {
    return message.success("Change body successfully!");
  }
  return message.error("Failed!");
};
const updateMockApi = async (response: object, path: string) => {
  const api = {
    collection_id: collection_parent.value.id,
    path: path,
    method: response.method,
    status_code: 200,
    response_body: response.body,
  };

  response_saving.value = true;

  loadingBar.start();
  disabledRef.value = false;
  let result;
  const method = response.method === "GET" ? "POST" : response.method;
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/mock-api/update",
    {
      method: method,
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(api),
    }
  );
  result = JSON.parse(JSON.stringify(status.value));

  if (result.status) {
    const { data: status } = await useFetch(
      "http://localhost:8000/admin/response.json",
      {
        method: "PATCH",
        headers: { "Content-type": "application/json" },
        body: JSON.stringify(response),
      }
    );
    result = JSON.parse(JSON.stringify(status.value));
    loadingBar.finish();
    disabledRef.value = true;
    response_saving.value = false;
    if (result.status) {
      return message.success("Change Mock API successfully!");
    }
  }
  return message.error("Failed!");
};
</script>
<style></style>
