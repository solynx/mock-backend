<template>
  <n-breadcrumb separator="|">
    <n-breadcrumb-item><n-icon :component="MdCash" /> Home</n-breadcrumb-item>
    <n-breadcrumb-item> Account</n-breadcrumb-item>
    <n-breadcrumb-item> Category</n-breadcrumb-item>
  </n-breadcrumb>
</template>
<script setup>
import { NBreadcrumb, NBreadcrumbItem } from "naive-ui";
</script>
