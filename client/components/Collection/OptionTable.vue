<template>
  <NMessageProvider>
    <n-tabs type="line" default-value="params" size="large">
      <NTabPane name="var" tab="Variables">
        <CollectionVariableTable></CollectionVariableTable>
      </NTabPane>
    </n-tabs>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import { NTabs, NTabPane, NMessageProvider } from "naive-ui";
</script>

<style lang="scss" scoped></style>
