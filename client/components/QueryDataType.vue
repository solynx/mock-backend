<template>
  <n-tabs type="line" default-value="params" size="large">
    <NTabPane name="params" tab="Params">
      <ParamsTable></ParamsTable>
    </NTabPane>

    <NTabPane name="params1" tab="Headers">
      <HeadersTable></HeadersTable>
    </NTabPane>
    <NTabPane name="params2">
      <template #tab>
        <span>Body</span>
        <span class="mx-2" v-if="request_selected.body">
          <n-icon size="8" color="#0cbb52"><RadioModifier /></n-icon
        ></span>
      </template>
      <QueryTypeY></QueryTypeY>
    </NTabPane>
  </n-tabs>
</template>

<script lang="ts" setup>
import { NTabs, NTab, NTabPane, NButton, NIcon } from "naive-ui";
import { RadioButtonOnOutline as RadioModifier } from "@vicons/ionicons5";

const request_selected = useState("request_selected");
</script>

<style lang="scss" scoped></style>
