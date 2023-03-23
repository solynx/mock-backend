<template>
  <n-tabs type="line" default-value="params" size="large">
    <NTabPane name="params" tab="Params">
      <ParamsTable></ParamsTable>
    </NTabPane>

    <NTabPane name="params1" tab="Headers">
      <HeadersTable></HeadersTable>
    </NTabPane>
    <NTabPane name="params2" tab="Body">
      <QueryTypeY></QueryTypeY>
    </NTabPane>
  </n-tabs>
</template>

<script setup>
import { NTabs, NTab, NTabPane, NButton } from "naive-ui";
</script>

<style lang="scss" scoped></style>
