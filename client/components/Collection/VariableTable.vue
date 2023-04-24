<template>
  <n-scrollbar style="max-height: 320px">
    <div class="w-full">
      <n-data-table
        @update:checked-row-keys="handleCheck"
        :columns="columns"
        :data="data"
        :max="5"
      />
    </div>
  </n-scrollbar>
  <n-button tertiary type="warning" class="float-right" @click="saveVariable">
    Save Variable
  </n-button>
</template>
<script lang="ts" setup>
import {
  NScrollbar,
  NTable,
  DataTableColumns,
  NInput,
  NDataTable,
  DataTableRowKey,
  useMessage,
  NButton,
} from "naive-ui";
type RowData = {
  key: string;
  param: string;
  value: string;
  description: string;
};
const message = useMessage();
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
const data = useState("var_collection", () => createData());
const var_actived = useState("var_actived", () => {});
const collection_variable = useState("collection_variable");
const emit = defineEmits(["binding_param", "active_query"]);
const createColumns = (): DataTableColumns<RowData> => [
  {
    title: "Variable",
    key: "abc1",
    render(row, index) {
      return h(NInput, {
        value: row.param,
        onUpdateValue(v) {
          data.value[index].param = v;
          emit("binding_param", data.value);
        },

        placeholder: "Variable",
      });
    },
  },
  {
    title: "Initial Value",
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
    title: "Current Value",
    key: "current",
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

const saveVariable = async () => {
  let arr_var = [];
  data.value.forEach((item) => {
    if (item.param === "" && item.value === "") return;
    let new_var = {
      key: item.param,
      value: item.value,
      type: "string",
    };
    arr_var.push(new_var);
  });

  const jsonString = JSON.stringify(arr_var);
  const base64String = btoa(jsonString);
  collection_variable.value.variable = base64String;

  const collection = {
    id: collection_variable.value.id,
    variable: base64String,
  };
  console.log(collection);
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/collection.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(collection),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));
  if (status1.status) {
    alert("Add variable to collection successfully!");
  } else {
    alert("Failed!");
  }
};
</script>
