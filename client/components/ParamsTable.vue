<template>
  <n-scrollbar style="max-height: 320px">
    <div class="query-group w-full">
      <n-data-table
        @update:checked-row-keys="handleCheck"
        :columns="columns"
        :data="data"
      />
    </div>
  </n-scrollbar>
</template>
<script lang="ts" setup>
import {
  NScrollbar,
  NTable,
  DataTableColumns,
  NInput,
  NDataTable,
  DataTableRowKey,
} from "naive-ui";
import { useUriQP } from "~~/composables/useUriQP";
type RowData = {
  key: string;
  param: string;
  value: string;
  description: string;
};
// const data = useState("data", () => createData());
// const same_data = useState("data");
const data = useUriQP();

const emit = defineEmits(["binding_param", "active_query"]);
const createColumns = (): DataTableColumns<RowData> => [
  {
    type: "selection",
    fixed: "left",
  },
  {
    title: "Key",
    key: "abc1",
    render(row, index) {
      return h(NInput, {
        value: row.param,
        onUpdateValue(v) {
          data.value[index].param = v;
          const test = useQueryTest();
        },

        placeholder: "Key",
      });
    },
  },
  {
    title: "Value",
    key: "abc1",
    render(row, index) {
      return h(NInput, {
        value: row.value,
        onUpdateValue(v) {
          data.value[index].value = v;
          emit("binding_param", data.value);
        },

        placeholder: "Value",
      });
    },
  },
  {
    title: "Description",
    key: "des",
    render(row, index) {
      return h(NInput, {
        value: row.description,
        onUpdateValue(v) {
          data.value[index].description = v;
        },
        placeholder: "Description",
      });
    },
  },
];
const columns = createColumns();
const checkedRowKeysRef = ref<Array<string | number>>([]);
const handleCheck = (rowKeys: DataTableRowKey[]) => {
  checkedRowKeysRef.value = rowKeys;

  emit("active_query", checkedRowKeysRef.value, data.value);
};
</script>
