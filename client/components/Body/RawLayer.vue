<template>
  <n-input
    class="mt-3"
    v-model:value="json_raw.body"
    type="textarea"
    placeholder=""
    rows="8"
  />
</template>

<script lang="ts" setup>
import { NSpace, NInput } from "naive-ui";
const json_raw = useState("request_selected");
</script>
