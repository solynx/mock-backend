<template>
  <NSpace>
    <NDropdown trigger="hover" :options="options" @select="handleSelect">
      <NButton>Hover!</NButton>
    </NDropdown>
    <NDropdown trigger="click" :options="options" @select="handleSelect">
      <NButton>Click!</NButton>
    </NDropdown>
    <NDropdown :show="showDropdown" :options="options" @select="handleSelect">
      <NButton @click="handleClick"> Oh! Manually By Myself! </NButton>
    </NDropdown>
  </NSpace>
</template>

<script lang="ts">
import {
  NMessageProvider,
  NNotificationProvider,
  NDialog,
  NLoadingBarProvider,
  NButton,
  NSpace,
  NDropdown,
} from "naive-ui";
import { defineComponent, ref } from "vue";

export default defineComponent({
  setup() {
    const showDropdownRef = ref(false);
    return {
      options: [
        {
          label: "Marina Bay Sands",
          key: "Marina Bay Sands",
          disabled: true,
        },
        {
          label: "Brown's Hotel, London",
          key: "Brown's Hotel, London",
        },
        {
          label: "Atlantis Bahamas, Nassau",
          key: "Atlantis Bahamas, Nassau",
        },
        {
          label: "The Beverly Hills Hotel, Los Angeles",
          key: "The Beverly Hills Hotel, Los Angeles",
        },
      ],
      showDropdown: showDropdownRef,
      handleSelect(key: string | number) {
        message.info(String(key));
      },
      handleClick() {
        showDropdownRef.value = !showDropdownRef.value;
      },
    };
  },
});
</script>
