<template>
  <n-card>
    <div class="w-full">
      <div class="text-step w-2/5">
        <n-gradient-text size="24" type="info">
          Create a mock server
        </n-gradient-text>
        <n-tabs
          default-value="select_collection"
          size="small"
          justify-content="space-evenly"
          :value="step_toggle"
        >
          <n-tab-pane name="select_collection" tab="a">
            <template #tab>
              <p>1. Select collection to mock</p>
            </template>
          </n-tab-pane>
          <n-tab-pane name="mock_name" tab="">
            <template #tab>
              <p>2. Configuration</p>
            </template>
          </n-tab-pane>
        </n-tabs>
      </div>
      <div v-if="step_num" class="step-1">
        <n-tag>Create a new collection</n-tag>
        <p class="pt-5">
          Enter the requests you want to mock. Optionally, add a request body by
          clicking on the (...) icon.
        </p>
        <div class="pt-5">
          <n-data-table :columns="columns" :data="data" />
        </div>

        <div class="float-right mt-5">
          <n-space>
            <n-button strong secondary type="tertiary"> Cancel </n-button>
            <n-button strong secondary type="info" @click="changeStep">
              Next
            </n-button>
          </n-space>
        </div>
      </div>
      <div v-else class="step-2 mt-5">
        <n-space vertical>
          <div class="grid grid-cols-6 gap-4">
            <div class="col-start-2 col-span-4">
              <n-tag>Mock server name</n-tag>
              <n-input
                v-model:value="collection_name"
                type="text"
                placeholder="Server name"
                class="mt-3"
              />
            </div>
          </div>
        </n-space>
        <div class="float-right mt-5">
          <n-button strong secondary type="tertiary" @click="changeStep">
            Cancel
          </n-button>
          <n-button strong secondary type="info" @click="createMockServer"
            >Create Mock Server</n-button
          >
        </div>
      </div>
    </div>
  </n-card>
</template>

<script lang="ts" setup>
import {
  NCard,
  NTabs,
  NTabPane,
  NIcon,
  NGradientText,
  NTag,
  NDataTable,
  DataTableColumns,
  NInput,
  NSelect,
  NButton,
  NSpace,
  useMessage,
} from "naive-ui";
import { CaretDownCircle as Step } from "@vicons/ionicons5";
import { result } from "lodash";
const message = useMessage();
const collection_name = ref("");
const step_toggle = ref("select_collection");
// const props = defineProps({
//   servers: Array,
// });
const emit = defineEmits(["createServer"]);
type RowData = {
  key: number;
  method: string;
  url: string;
  response_code: number;
  response_body: string;
};
const createData = (): RowData[] => [
  {
    key: 0,
    method: "GET",
    url: "",
    response_code: 200,
    response_body: "",
  },
];
const options = [
  {
    label: "GET",
    value: "GET",
  },
  {
    label: "POST",
    value: "POST",
  },
  {
    label: "PATCH",
    value: "PATCH",
  },
  {
    label: "DELETE",
    value: "DELETE",
  },
];
const data = ref(createData());
const createColumns = (): DataTableColumns<RowData> => [
  {
    title: "Request Method",
    key: "method",
    render(row, index) {
      return h(NSelect, {
        value: row.method,
        options: options,
        onUpdateValue(v) {
          data.value[index].method = v;
        },
      });
    },
  },
  {
    title: "Request URL",
    key: "url",
    render(row, index) {
      return h(NInput, {
        value: row.url,
        onUpdateValue(v) {
          data.value[index].url = v;
        },
      });
    },
  },
  {
    title: "Response Code",
    key: "status_code",
    render(row, index) {
      return h(NInput, {
        value: row.response_code,
        onUpdateValue(v) {
          data.value[index].response_code = v;
        },
      });
    },
  },
  {
    title: "Response Body",
    key: "body",
    render(row, index) {
      return h(NInput, {
        value: row.response_body,
        onUpdateValue(v) {
          data.value[index].response_body = v;
        },
      });
    },
  },
];
const columns = createColumns();
const step_num = ref(true);
const changeStep = () => {
  step_num.value = !step_num.value;
  step_toggle.value =
    step_toggle.value == "mock_name" ? "select_collection" : "mock_name";
};
const createMockServer = () => {
  if (/^localhost:8000\//.test(data.value[0].url)) {
    const url = new URL("http://" + data.value[0].url);
    const path = url.pathname;
    data.value[0].url = path;

    emit("createServer", data.value[0], collection_name.value);
  }
};
</script>

<style></style>
