<template>
  <BodyNoneLayer v-if="props.layer === 'none'" />
  <BodyRawLayer v-if="props.layer === 'raw'" />
  <BodyBinaryLayer v-if="props.layer === 'binary'" />
</template>

<script lang="ts" setup>
const props = defineProps({
  layer: String,
});
</script>

<style lang="scss" scoped></style>
