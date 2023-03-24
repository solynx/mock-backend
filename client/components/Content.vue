<template>
  <n-card>
    <div class="w-full">
      <!-- Phần tag multiple bao gồm có nhiều file được mở cùng lúc tại  thanh này-->
      <div class="file-multiple">
        <n-dynamic-tags v-model:value="tags" />
      </div>
      <n-card>
        <div class="file-detail flex justify-between">
          <div class="file-link">
            <FileLink :link="item_link"></FileLink>
          </div>
          <FileOption></FileOption>
        </div>
      </n-card>

      <URIComponent
        :uri="item"
        @change_link="getRequest"
        :filter_param="filteredQueryParam.value"
      ></URIComponent>

      <div class="group-func flex">
        <n-card style="width: 60%; min-width: 50%">
          <div class="group-fun-query">
            <div class="request-option">
              <QueryDataType></QueryDataType>
            </div>
          </div>
        </n-card>

        <n-config-provider style="width: 40%; max-width: 50%" :hljs="hljs">
          <n-loading-bar-provider>
            <NSpin :show="!disabledRef" size="small">
              <ResponseData></ResponseData>
              <template #description> Just moment... </template>
            </NSpin>
          </n-loading-bar-provider>
        </n-config-provider>
      </div>
    </div></n-card
  >
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
} from "naive-ui";
import hljs from "highlight.js/lib/core";
import json from "highlight.js/lib/languages/json";
hljs.registerLanguage("json", json);

const props = defineProps({
  item: Object,
  item_link: Array,
});
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

const code = useState("code_response", () => "");
const getRequest = async (request_detail: object) => {
  loadingBar.start();
  disabledRef.value = false;
  let result;
  const { data: status } = await useFetch(request_detail.link, {
    method: request_detail.method,
    headers: { "Content-type": "application/json" },
  });
  result = JSON.parse(JSON.stringify(status.value));
  loadingBar.finish();
  disabledRef.value = true;
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
const tags = ref(["teacher", "programmer"]);
const checkedValueRef = ref<string | null>(null);
const checkedValue = checkedValueRef;
</script>
