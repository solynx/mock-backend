<template>
  <NMessageProvider>
    <NDialogProvider>
      <NuxtLayout name="sidebar"></NuxtLayout>
    </NDialogProvider>
  </NMessageProvider>
</template>

<script lang="ts" setup>
import { NMessageProvider, NDialogProvider } from "naive-ui";
</script>
<style scoped>
html {
  overflow-x: hidden;
}
</style>
