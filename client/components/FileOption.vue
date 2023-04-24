<template>
  <div class="file-option">
    <n-config-provider>
      <n-space>
        <n-button-group>
          <n-button
            tertiary
            :class="{ saving: response_saving }"
            @click="CheckMockApi"
          >
            <n-icon size="24">
              <save-icon />
            </n-icon>
            <b> {{ response_saving ? "Saving..." : "Save" }}</b>
          </n-button>
          <n-button tertiary>
            <n-icon>
              <chevron-down-icon />
            </n-icon>
          </n-button>
          <n-button tertiary>
            <n-icon size="24">
              <more-icon />
            </n-icon>
          </n-button>
        </n-button-group>
        <n-tag :bordered="false" class="w-full h-full">
          <n-icon size="20">
            <pencil-icon class="pencil-icon-background" />
          </n-icon>

          <n-button text>
            <n-icon size="20">
              <moment-mode-icon />
            </n-icon>
          </n-button>
        </n-tag>
      </n-space>
    </n-config-provider>
  </div>
</template>
<script lang="ts" setup>
import {
  NButton,
  NSpace,
  NConfigProvider,
  NIcon,
  NButtonGroup,
  NTag,
} from "naive-ui";
import {
  SaveOutline as SaveIcon,
  ChevronDownOutline as ChevronDownIcon,
  EllipsisHorizontalOutline as MoreIcon,
  PencilOutline as PencilIcon,
  ChatboxEllipsesOutline as MomentModeIcon,
} from "@vicons/ionicons5";
const emit = defineEmits(["update_response", "update_api"]);
const response_saving = useState("response_saving");
const response_actived = useState("response_actived");
const collection_parent = useState("belong_collection");
const request_parent = useState("belong_request");
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
<style scoped>
.pencil-icon-background {
  background-color: white;
}
.fixed-background-moment-icon {
  background-color: rgb(212 212 212);
}
.saving {
  opacity: 0.5;
}
</style>
