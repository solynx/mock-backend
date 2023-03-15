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
            <FileLink></FileLink>
          </div>
          <FileOption></FileOption>
        </div>
      </n-card>
      <URIComponent></URIComponent>
      <div class="request-option">
        <QueryDataType></QueryDataType>
        <QueryTypeY></QueryTypeY>
      </div>

      <div class="group-func flex">
        <n-card>
          <div class="group-fun-query">
            <n-scrollbar style="max-height: 320px">
              <div class="query-group w-full">
                <NTable :bordered="false" :single-line="false">
                  <thead>
                    <tr>
                      <th></th>
                      <th>Key</th>
                      <th>Value</th>
                      <th>Description</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td></td>
                      <td>...</td>
                      <td>Damn it! I</td>
                      <td>...</td>
                    </tr>
                    <tr>
                      <td></td>
                      <td>...</td>
                      <td>...</td>
                      <td>...</td>
                    </tr>
                  </tbody>
                </NTable>
              </div>
            </n-scrollbar>
          </div>
        </n-card>
        <n-scrollbar style="max-height: 320px">
          <div class="json-result w-full">
            <NConfigProvider :hljs="hljs">
              <div style="overflow: auto">
                <n-code :code="code" language="cpp" show-line-numbers />
              </div>
            </NConfigProvider>
          </div>
        </n-scrollbar>
      </div></div
  ></n-card>
</template>

<script lang="ts" setup>
import { defineComponent, ref } from "vue";
import hljs from "highlight.js/lib/core";
import javascript from "highlight.js/lib/languages/javascript";
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
} from "naive-ui";

hljs.registerLanguage("javascript", javascript);

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

const code = `#include <bits/stdc++.h>
using namespace std;

int main() {
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  cout <<"\\n" <<endl;
  return 0;
}`;

const tags = ref(["teacher", "programmer"]);
const checkedValueRef = ref<string | null>(null);
const checkedValue = checkedValueRef;
</script>
