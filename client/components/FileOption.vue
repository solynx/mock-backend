<template>
  <div class="file-option">
    <n-config-provider>
      <n-space>
        <n-button-group>
          <n-button tertiary>
            <n-icon size="24">
              <save-icon />
            </n-icon>
            <b> Save</b>
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
<script setup>
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
</script>
<style scoped>
.pencil-icon-background {
  background-color: white;
}
.fixed-background-moment-icon {
  background-color: rgb(212 212 212);
}
</style>
