<template>
  <n-scrollbar style="max-height: 400px">
    <NSpace vertical>
      <n-input
        v-model:value="response_actived.body"
        placeholder="Autosizable"
        type="textarea"
        size="small"
        :autosize="{
          minRows: 16,
          maxRows: 100,
        }"
        @click.ctrl="CheckMockApi()"
      />
    </NSpace>
  </n-scrollbar>
</template>
<script lang="ts" setup>
import { NSpace, NInput, NScrollbar } from "naive-ui";
const emit = defineEmits(["update_response", "update_api"]);
const response_actived = useState("response_actived");
const collection_parent = useState("belong_collection");
const request_parent = useState("belong_request");
const response_saving = useState("response_saving", () => false);

const CheckMockApi = () => {
  if (collection_parent.value.is_server) {
    const url = new URL(response_actived.value.uri_component);
    const path = url.pathname;
    const full_path = path.split("/");
    const uuid = full_path[1];
    const after_uuid_path = "/" + full_path.slice(2).join("/");
    if (uuid === collection_parent.value.id) {
      return emit("update_api", response_actived.value, after_uuid_path);
    }
  }
  return emit("update_response", response_actived.value);
};
</script>
