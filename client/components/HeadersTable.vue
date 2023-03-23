<template>
  <n-scrollbar style="max-height: 320px">
    <div class="query-group w-full">
      <n-data-table
        @update:checked-row-keys="handleCheck"
        :columns="columns"
        :data="data1"
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
type RowData = {
  key: string;
  param: string;
  value: string;
  description: string;
};
const createData = (): RowData[] => [
  {
    key: "aaaa",
    param: "",
    value: "",
    description: "",
  },
  {
    key: "a11",
    param: "",
    value: "",
    description: "",
  },
];
// const data = useState("data", () => createData());
// const same_data = useState("data");
const data1 = createData();

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
          emit("binding_param", data.value);
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
};
</script>
