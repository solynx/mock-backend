<template>
  <n-scrollbar style="max-height: 320px">
    <div class="query-group w-full">
      <n-data-table
        @update:checked-row-keys="handleCheck"
        :columns="columns"
        :data="data"
        :default-checked-row-keys="checkedRowKeys1"
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
  key: number;
  param: string;
  value: string;
  description: string;
};
// const data = useState("data", () => createData());
// const same_data = useState("data");

const data = useUriQP();
const query = useState("query_active");
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
          if (data.value.length === row.key && data.value.length < 5) {
            const newRow: RowData = {
              key: data.value.length + 1,
              param: "",
              value: "",
              description: "",
            };
            data.value.push(newRow);
          }
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
const checkedRowKeysRef = ref<Array<string | number>>([1]);
const checkedRowKeys1 = useState("checkedRowKeys1");
const handleCheck = (rowKeys: DataTableRowKey[]) => {
  checkedRowKeysRef.value = rowKeys;
  const newArr: Array<object> = [];
  checkedRowKeysRef.value.forEach((key) => {
    return newArr.push(data.value[key - 1]);
  });

  query.value = newArr;

  // uriStore.queries = newArr;
  // uriStore.convertUrl();
  emit("active_query", checkedRowKeysRef.value, data.value);
};
</script>
