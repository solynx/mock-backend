<template>
  <n-scrollbar style="max-height: 320px">
    <div class="query-group w-full">
      <n-data-table
        @update:checked-row-keys="handleCheck"
        :columns="columns"
        :data="data"
        :max="5"
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
  NButton,
} from "naive-ui";
type RowData = {
  key: string;
  param: string;
  value: string;
  description: string;
};
const createData = (): RowData[] => [
  {
    key: 1,
    param: "",
    value: "",
    description: "",
  },
];
// const data = useState("data", () => createData());
// const same_data = useState("data");
const data = useState("header_req", () => createData());
const header_actived = useState("header_actived", () => {});

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
const checkedRowKeysRef = ref<Array<string | number>>([]);
const handleCheck = (rowKeys: DataTableRowKey[]) => {
  checkedRowKeysRef.value = rowKeys;
  const newArr: Array<object> = [];
  checkedRowKeysRef.value.forEach((key) => {
    return newArr.push(data.value[key - 1]);
  });
  header_actived.value = [];
  newArr.forEach((item) => {
    let key = item.param;
    let value = item.value;
    let header = { [key]: value };

    header_actived.value[key] = value;
  });

  // uriStore.queries = newArr;
  // uriStore.convertUrl();
};
</script>
