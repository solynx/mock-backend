<template>
  <n-scrollbar style="max-height: 400px">
    <div class="json-result w-full">
      <div style="overflow: auto" class="px-12">
        <n-code :code="code_res" language="json" show-line-numbers />
      </div>
    </div>
  </n-scrollbar>
</template>
<script lang="ts" setup>
import { NScrollbar, NCode } from "naive-ui";

// const props = defineProps({
//   code: String,
// });
const code_res = useState("code_response");
</script>
