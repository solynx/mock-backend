<template>
  <n-breadcrumb separator="/">
    <n-breadcrumb-item v-for="item in link"> {{ item.name }}</n-breadcrumb-item>
  </n-breadcrumb>
</template>
<script setup>
import { NBreadcrumb, NBreadcrumbItem } from "naive-ui";
const props = defineProps({
  link: Object,
});
</script>
