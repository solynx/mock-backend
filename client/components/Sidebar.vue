<template>
  <div class="sidebar h-screen bg-neutral-300">
    <nav class="sidebar__group">
      <ul class="sidebar__group__list flex py-5">
        <li class="sidebar__group_list__item sidebar__group_list__item--create">
          <n-button
            v-if="
              toggle_mock_collection !== 'mock_server' &&
              toggle_mock_collection !== 'show_create_mock'
            "
            size="small"
            tertiary
            @click="newCollection"
          >
            <n-icon class="sidebar__group__list--icon text-2xl">
              <add-icon></add-icon>
            </n-icon>
          </n-button>
          <n-button
            v-else
            size="small"
            tertiary
            @click="toggleMock('show_create_mock')"
          >
            <n-icon class="sidebar__group__list--icon text-2xl">
              <add-icon></add-icon>
            </n-icon>
          </n-button>
        </li>
        <li class="sidebar__group__list__item">
          <n-input placeholder="Tìm kiếm" size="small">
            <template #prefix>
              <n-icon :component="FilterIcon" />
            </template>
          </n-input>
        </li>
        <li class="sidebar__group__list__item">
          <NDropdown trigger="click" :options="NearFilterOptions">
            <NButton size="small" tertiary>
              <n-icon size="20">
                <more-icon />
              </n-icon>
            </NButton>
          </NDropdown>
        </li>
      </ul>
      <n-upload
        @finish="handleFinish"
        action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
        ref="uploadRef"
      >
        <n-button size="small" :default-upload="false" tertiary
          >Import</n-button
        >
      </n-upload>

      <NScrollbar style="max-height: 85vh">
        <n-tabs
          default-value="collection"
          size="small"
          justify-content="space-evenly"
        >
          <n-tab-pane name="collection">
            <template #tab>
              <p @click="toggleMock('collection')">Collections</p>
            </template>
            <ul class="sidebar__group__collection">
              <n-collapse :arrow="false" :accordion="true">
                <n-collapse-item
                  v-for="(collection, index) in props.collections"
                  :title="`${collection.name}`"
                  :name="collection.id"
                  :key="collection.id"
                >
                  <template #header>
                    <span v-if="editingIndex !== collection.id">
                      {{ collection.name }}</span
                    >

                    <input
                      v-else
                      type="text"
                      v-model="name"
                      @blur="handleBlur"
                      @keyup.enter="saveNameCollection(collection, name)"
                    />
                    <n-button
                      size="tiny"
                      secondary
                      type="tertiary"
                      @click="exportModal(collection, true)"
                    >
                      <small> Export</small>
                    </n-button>
                    <n-modal
                      v-model:show="mode"
                      class="custom-card"
                      preset="card"
                      :style="{ width: 600 + 'px' }"
                      title="Export Collection"
                      :bordered="false"
                      size="huge"
                      :segmented="segmented"
                    >
                      <div>
                        <b>{{ collection.name }}</b>
                      </div>
                      <div>will be exported as a JSON file. Export as:</div>
                      <n-radio-group
                        v-model:value="json_version"
                        name="radiogroup"
                      >
                        <n-space>
                          <n-radio
                            key="Collectionv2"
                            value="Collectionv2"
                            label="Collection v2"
                          />
                          <br />
                          <n-radio
                            key="Collectionv2.1"
                            value="Collectionv2.1"
                            label="Collection v2.1 (recommended)"
                          />
                        </n-space>
                      </n-radio-group>
                      <template #footer>
                        <n-space class="float-right">
                          <n-button
                            strong
                            secondary
                            type="tertiary"
                            @click="exportModal(collection, false)"
                          >
                            Cancel
                          </n-button>
                          <n-button
                            strong
                            class="no-tailwind"
                            color="#ff6c37"
                            @click="dowloadFile"
                          >
                            Export
                          </n-button>
                        </n-space>
                      </template>
                    </n-modal>
                  </template>

                  <template #header-extra>
                    <n-space>
                      <n-popover trigger="click">
                        <template #trigger>
                          <n-button>
                            <n-icon> <more-icon></more-icon> </n-icon
                          ></n-button>
                        </template>
                        <template #header>
                          <n-button @click="addRequest(collection)">
                            Add Request
                          </n-button>
                        </template>
                        <div>
                          <n-button @click="newFolder(collection)">
                            Add Folder
                          </n-button>
                        </div>
                        <n-button @click="toggleEdit1(collection)">
                          Edit
                        </n-button>
                        <template #footer>
                          <n-button
                            @click="deleteCollection(collection, index)"
                          >
                            Delete
                          </n-button>
                        </template>
                      </n-popover>
                    </n-space>
                  </template>
                  <!-- FOLDERS IN COLLECTION-->
                  <n-collapse>
                    <n-collapse-item
                      v-for="(folder, index) in collection.folders"
                      :title="`${folder.name}`"
                      :name="folder.name"
                    >
                      <template #header>
                        <span v-if="editingIndex !== folder.id">
                          {{ folder.name }}</span
                        >
                        <input
                          v-else
                          type="text"
                          v-model="name"
                          @blur="handleBlur"
                          @keyup.enter="saveNameFolder(folder, name)"
                        />
                      </template>

                      <template #header-extra>
                        <n-space>
                          <n-popover trigger="click">
                            <template #trigger>
                              <n-button>
                                <n-icon> <more-icon></more-icon> </n-icon
                              ></n-button>
                            </template>
                            <template #header>
                              <n-button @click="addRequest(collection, folder)">
                                Add Request
                              </n-button>
                            </template>
                            <n-button @click="toggleEdit1(folder)">
                              Edit
                            </n-button>
                            <template #footer>
                              <n-button
                                @click="deleteFolder(collection, folder, index)"
                              >
                                Delete
                              </n-button>
                            </template>
                          </n-popover>
                        </n-space>
                      </template>
                      <n-collapse>
                        <!-- Request In Folder-->
                        <n-collapse v-for="(request, index) in folder.requests">
                          <n-collapse-item
                            :title="`${request.name}`"
                            :name="request.id"
                          >
                            <template #header>
                              <span
                                v-if="editingIndex !== request.id"
                                @click="
                                  getItemSelected(
                                    [
                                      {
                                        id: collection.id,
                                        name: collection.name,
                                      },
                                      { id: folder.id, name: folder.name },
                                    ],
                                    request
                                  )
                                "
                              >
                                {{ request.name }}</span
                              >
                              <input
                                v-else
                                type="text"
                                v-model="name"
                                @blur="handleBlur"
                                @keyup.enter="saveNameRequest(request, name)"
                              />
                            </template>
                            <template #header-extra>
                              <n-space>
                                <n-popover trigger="click">
                                  <template #trigger>
                                    <n-button>
                                      <n-icon> <more-icon></more-icon> </n-icon
                                    ></n-button>
                                  </template>
                                  <template #header>
                                    <n-button @click="addResponse(request)">
                                      Add Response
                                    </n-button>
                                  </template>
                                  <n-button @click="toggleEdit1(request)">
                                    Edit
                                  </n-button>
                                  <template #footer>
                                    <n-button
                                      @click="
                                        deleteRequest(folder, request, index)
                                      "
                                    >
                                      Delete
                                    </n-button>
                                  </template>
                                </n-popover>
                              </n-space>
                            </template>
                            <!-- RESPONSE IN REQUEST OF folders-->

                            <n-collapse-item
                              v-for="(res, index) in request.responses"
                              :title="`${res.name}`"
                              :name="res.id"
                            >
                              <template #header>
                                <span
                                  v-if="editingIndex !== res.id"
                                  @click="
                                    getResponseSelected(
                                      collection,
                                      folder,
                                      request,
                                      res
                                    )
                                  "
                                >
                                  {{ res.name }}</span
                                >
                                <input
                                  v-else
                                  type="text"
                                  v-model="name"
                                  @blur="handleBlur"
                                  @keyup.enter="editResponse(res, name)"
                                />
                              </template>

                              <template #header-extra>
                                <n-space>
                                  <n-popover trigger="click">
                                    <template #trigger>
                                      <n-button>
                                        <n-icon>
                                          <more-icon></more-icon> </n-icon
                                      ></n-button>
                                    </template>
                                    <template #header>
                                      <n-button @click="toggleEdit1(res)">
                                        Edit
                                      </n-button>
                                    </template>
                                    <template #footer>
                                      <n-button
                                        @click="
                                          deleteResponse(request, res, index)
                                        "
                                      >
                                        Delete
                                      </n-button>
                                    </template>
                                  </n-popover>
                                </n-space>
                              </template>
                            </n-collapse-item>
                          </n-collapse-item>
                        </n-collapse>
                      </n-collapse>
                    </n-collapse-item>
                  </n-collapse>
                  <!-- REQUEST IN COLLECTION-->
                  <n-collapse>
                    <n-collapse-item
                      v-for="(request, index) in collection.requests"
                      :title="`${request.name}`"
                      :name="request.id"
                    >
                      <template #header>
                        <span
                          v-if="editingIndex !== request.id"
                          @click="
                            getItemSelected(
                              [{ id: collection.id, name: collection.name }],
                              request
                            )
                          "
                        >
                          {{ request.name }}</span
                        >
                        <input
                          v-else
                          type="text"
                          v-model="name"
                          @blur="handleBlur"
                          @keyup.enter="saveNameRequest(request, name)"
                        />
                      </template>

                      <template #header-extra>
                        <n-space>
                          <n-popover trigger="click">
                            <template #trigger>
                              <n-button>
                                <n-icon> <more-icon></more-icon> </n-icon
                              ></n-button>
                            </template>
                            <template #header>
                              <n-button @click="addResponse(request)">
                                Add Response
                              </n-button>
                            </template>

                            <n-button @click="toggleEdit1(request)">
                              Edit
                            </n-button>
                            <template #footer>
                              <n-button
                                @click="
                                  deleteRequest(collection, request, index)
                                "
                              >
                                Delete
                              </n-button>
                            </template>
                          </n-popover>
                        </n-space>
                      </template>
                      <!-- RESPONSE IN REQUEST OF COLLECTIONs-->

                      <n-collapse-item
                        v-for="(res, index) in request.responses"
                        :title="`${res.name}`"
                        :name="res.id"
                      >
                        <template #header>
                          <span
                            v-if="editingIndex !== res.id"
                            @click="
                              getResponseSelected(
                                collection,
                                null,
                                request,
                                res
                              )
                            "
                          >
                            {{ res.name }}</span
                          >
                          <input
                            v-else
                            type="text"
                            v-model="name"
                            @blur="handleBlur"
                            @keyup.enter="
                              saveName(request, res, name, 'RESPONSE')
                            "
                          />
                        </template>

                        <template #header-extra>
                          <n-space>
                            <n-popover trigger="click">
                              <template #trigger>
                                <n-button>
                                  <n-icon> <more-icon></more-icon> </n-icon
                                ></n-button>
                              </template>
                              <template #header>
                                <n-button @click="toggleEdit1(res)">
                                  Edit
                                </n-button>
                              </template>
                              <template #footer>
                                <n-button
                                  @click="deleteResponse(request, res, index)"
                                >
                                  Delete
                                </n-button>
                              </template>
                            </n-popover>
                          </n-space>
                        </template>
                      </n-collapse-item>
                    </n-collapse-item>
                  </n-collapse>
                </n-collapse-item>
              </n-collapse>
            </ul>
          </n-tab-pane>
          <n-tab-pane name="mock" tab="Mock Servers">
            <template #tab>
              <n-icon size="14"> <server></server> </n-icon>
              <p @click="toggleMock('mock_server')">Mock Servers</p>
            </template>
            <n-collapse>
              <n-collapse-item
                v-for="server in props.servers"
                :title="server.name"
                name="server.name"
              >
              </n-collapse-item>
            </n-collapse>
          </n-tab-pane>
        </n-tabs>
      </NScrollbar>
    </nav>
  </div>
</template>
<script lang="ts" setup>
import {
  NInput,
  NIcon,
  NDropdown,
  NInputGroup,
  NInputGroupLabel,
  NCollapse,
  NCollapseItem,
  NButton,
  useMessage,
  NMessageProvider,
  NPopover,
  InputInst,
  NScrollbar,
  NModal,
  NRadioGroup,
  NRadio,
  NTabs,
  NTabPane,
  useNotification,
  NUpload,
  UploadFileInfo,
  useDialog,
  UploadInst,
} from "naive-ui";
import {
  AddOutline as AddIcon,
  FilterOutline as FilterIcon,
  EllipsisHorizontalOutline as MoreIcon,
  ServerOutline as Server,
} from "@vicons/ionicons5";
import { v4 as uuidv4 } from "uuid";
const props = defineProps({
  collections: Array,
  servers: Array,
});

const ModeAddButton = ref(true);
const item_selected = useState("item_select", () => "");
const response_actived = useState("response_actived", () => "");
const belong_collection = useState("belong_collection", () => "");
const belong_folder = useState("belong_folder", () => "");
const belong_request = useState("belong_request", () => "");
const notification = useNotification();
const dialog = useDialog();
const changeAddButton = () => {
  ModeAddButton.value = !ModeAddButton.value;
};
const uploadRef = ref<UploadInst | null>(null);
const emit = defineEmits(["item-selected", "toggle-layout"]);
const message = useMessage();
const treeOptions = [
  {
    label: "Add Folder",
    id: "key1",
  },
  {
    label: "Add Request",
    id: "key1",
  },
  {
    label: "Edit",
    id: "key2",
  },
  {
    label: "Delete",
    id: "key2",
  },
];

const options = [
  {
    label: "Marina Bay Sands",
    key: "Marina Bay Sands",
  },
];
const choose_collection = ref({});
// const toggle = ref(false);
// const toggleButton = (status: any) => {
//   toggle.value = !status.value;
//   console.log(toggle.value);
// };
const editingIndex = ref("a");
const inputInstRef = ref<InputInst | null>(null);
//Create New Collection
// const collections12 = ref(collections);
const mode = ref(false);
const toggle_mock_collection = useState(
  "toggle_mock_collection",
  () => "collection"
);
const newCollection = async () => {
  const uuid = uuidv4();
  const collection = {
    name: "New Collection",
    id: uuid,
    folders: [],
    requests: [],
  };

  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/collection.json",
    {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(collection),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    props.collections.push(collection);
    return message.success("Collection created successfully!");
  }

  return message.error("Collection created failed!");
};
const deleteCollection = async (collection: object, index: number) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/collection.json",
    {
      method: "DELETE",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(collection),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    props.collections.splice(index, 1);
    return message.success("Collection deleted successfully!");
  }

  return message.error("Collection deleted failed!");
};
const saveNameCollection = async (el: object, new_name: string) => {
  const collection = {
    id: el.id,
    name: new_name,
  };
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
    el.name = new_name;
    editingIndex.value = -1;
    return message.success("Collection created successfully!");
  }
  editingIndex.value = -1;
  return message.error("Collection created failed!");
};
const toggleEdit = true;
const optionName = ref(null);
const json_version = ref("Collectionv2.1");
let name = "A";
//---- MORE ICON
const handleSelect = (option: object) => {
  message.info(option);
};
const handleSelect1 = (key: string) => {
  console.log(key);
  message.info(key);
};
const handleCustomSelect = (option: any) => {
  console.log("Selected:", option);
  // Your custom logic here
};
//Request
const addRequest = async (collection: object, folder?: object) => {
  console.log(collection);
  const uuid = uuidv4();
  if (folder) {
    const request1 = {
      id: uuid,
      name: "New Request",
      folder_id: folder.id,
      collection_id: collection.id,
      uri_component: "",
      method: "GET",
      responses: [],
    };
    const result: object = await addRequestFetch(request1);

    if (result.status) {
      folder.requests.push(request1);
      return message.success("Success!");
    }
    return message.error("Error!");
  } else {
    const request = {
      id: uuid,
      name: "New Request",
      folder_id: null,
      collection_id: collection.id,
      uri_component: "",
      method: "GET",
      responses: [],
    };
    const result: object = await addRequestFetch(request);

    if (result.status) {
      collection.requests.push(request);
      return message.success("Success!");
    }
    return message.error("Error!");
  }
};

const addRequestFetch = async (request: object) => {
  let result;
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/request.json",
    {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(request),
    }
  );
  result = JSON.parse(JSON.stringify(status.value));

  return result;
};
const toggleEdit1 = (element: object) => {
  name = element.name;

  editingIndex.value = element.id;
};

const saveName = (
  parent_el?: object,
  el: object,
  new_name: string,
  type: string
) => {
  editingIndex.value = -1;
  el.name = new_name;
  name = null;
  console.log(new_name);
};

const handleBlur = () => {
  editingIndex.value = -1;
};
//FOLDER
const saveNameFolder = async (oldFolder: Object, new_name: string) => {
  const folder = {
    id: oldFolder.id,
    name: new_name,
  };
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/folder.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(folder),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));
  if (status1.status) {
    oldFolder.name = new_name;
    editingIndex.value = -1;
    return message.success("Success");
  }
  editingIndex.value = -1;
  return message.error("Failed!");
};
const deleteFolder = async (
  collection: object,
  folder: object,
  index: number
) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/folder.json",
    {
      method: "DELETE",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(folder),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));
  if (status1.status) {
    collection.folders.splice(index, 1);

    return message.success("Success!");
  }

  return message.error("Folder deleted failed!");
};
const newFolder = async (collection: object) => {
  const uuid = uuidv4();
  const folder = {
    name: "New Folder",
    id: uuid,
    collection_id: collection.id,
    requests: [],
  };
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/folder.json",
    {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(folder),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    collection.folders.push(folder);
    return message.success("Success");
  }
  return message.error("Failed!");
};

//Request
const saveNameRequest = async (request: object, new_name: string) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/request.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(request),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    request.name = new_name;
    editingIndex.value = -1;
    return message.success("Success");
  }
  editingIndex.value = -1;
  return message.error("Failed!");
};

const deleteRequest = async (
  parent_el: object,
  request: object,
  index: number
) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/request.json",
    {
      method: "DELETE",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(request),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));
  if (status1.status) {
    parent_el.requests.splice(index, 1);

    return message.success("Success");
  }

  return message.error("Failed!");
};
const addResponse = async (request: object) => {
  const uuid = uuidv4();
  const response = {
    id: uuid,
    name: "New response",
    method: request.method,
    uri_component: request.uri_component,
    request_id: request.id,
    file_link: null,
  };

  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/response.json",
    {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(response),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    request.responses.push(response);
    return message.success("Success");
  }
  return message.error("Failed!");
};
const deleteResponse = async (
  request: object,
  response: object,
  index: number
) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/response.json",
    {
      method: "DELETE",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(response),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    request.responses.splice(index, 1);
    return message.success("Response deleted successfully!");
  }
  return message.error("Failed!");
};
const editResponse = async (response: object, new_name: string) => {
  const { data: status } = await useFetch(
    "http://127.0.0.1:8000/admin/response.json",
    {
      method: "PATCH",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify(response),
    }
  );
  const status1 = JSON.parse(JSON.stringify(status.value));

  if (status1.status) {
    response.name = new_name;
    editingIndex.value = -1;
    return message.success("Rename successfully!");
  }
  editingIndex.value = -1;
  return message.error("Failed!");
};
// content action
const getItemSelected = (bread_cum: Array, item: object) => {
  item_selected.value = "request";

  emit("item-selected", bread_cum, item);
};
const getResponseSelected = (
  collection: object,
  folder: object,
  request: object,
  item: object
) => {
  const bread_cum =
    folder !== null ? [collection, folder, request] : [collection, request];

  item_selected.value = "response";
  response_actived.value = item;
  belong_collection.value = collection;
  belong_folder.value = folder !== null ? folder : "";
  belong_request.value = request;
  emit("item-selected", bread_cum, item);
};
const Alo = () => {
  console.log("alo");
};

const exportModal = (collection: object, choose: boolean) => {
  if (choose) {
    choose_collection.value = collection;
  }
  mode.value = !mode.value;
};
const newMockServer = () => {
  emit("toggle-layout");
};
const toggleMock = (element: string) => {
  toggle_mock_collection.value = element;
};
const downloadJSON = (data: object, name: string) => {
  // Tạo một đối tượng Blob từ dữ liệu JSON
  const blob = new Blob([JSON.stringify(data, null, 4)], {
    type: "application/json",
  });
  // Tạo một URL cho đối tượng Blob
  const url = URL.createObjectURL(blob);
  // Tạo một thẻ <a> ẩn với thuộc tính download và href là URL
  const link = document.createElement("a");
  link.download = name + ".collection_data.json";
  link.href = url;
  document.body.appendChild(link);
  link.click();
  // Xóa thẻ <a> và thu hồi URL
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
};
const dowloadFile = async () => {
  const { data } = await useFetch(
    "http://127.0.0.1:8000/admin/format-collection.json",
    {
      method: "POST",

      headers: { "Content-type": "application/json" },
      body: JSON.stringify(choose_collection.value),
    }
  );

  if (data.value != null) {
    downloadJSON(data.value, choose_collection.value.name);
    notification.success({
      content: "Your collection exported successfuly!",
      duration: 5000,
      closable: false,
      placement: "bottom-right",
    });
  }
  mode.value = !mode.value;
};
const handleFinish = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  if (file.type !== "application/json") {
    return message.error("File type must be .json");
  }
  let text = file.file;
  text?.text().then((text) => {
    let collection_import = JSON.parse(text);

    let check = props.collections?.find((collection) => {
      return collection.name === collection_import.info.name;
    });
    if (check != null) {
      dialog.warning({
        title: "Collection Exists",
        content: `A collection ${check.name} already exists. What would you like to do?"`,
        positiveText: "Import As Copy",
        negativeText: "Replace",
        onPositiveClick: () => {
          collection_import.info.name =
            collection_import.info.name +
            "- copy" +
            Math.floor(Math.random() * 100) +
            1;

          return addCollectionWithJsonFile(collection_import);
        },
        onNegativeClick: () => {
          message.error("Not Sure");
        },
      });
    } else {
      return addCollectionWithJsonFile(collection_import);
    }
  });
};
const addCollectionWithJsonFile = async (collection_import: object) => {
  const collection_copy = {
    id: uuidv4(),
    name: "",
    folders: [],
    requests: [],
  };
  collection_copy.name = collection_import.info.name;

  await useFetch("http://127.0.0.1:8000/admin/collection.json", {
    method: "POST",

    headers: { "Content-type": "application/json" },
    body: JSON.stringify(collection_copy),
  });
  if (collection_import.item !== null) {
    collection_import.item.forEach(async (item) => {
      if (item.item) {
        const folder = {
          id: uuidv4(),
          name: item.name,
          requests: [],
          collection_id: collection_copy.id,
        };
        const { data } = await useFetch(
          "http://127.0.0.1:8000/admin/folder.json",
          {
            method: "POST",

            headers: { "Content-type": "application/json" },
            body: JSON.stringify(folder),
          }
        );
        let requests = item.item; //arr
        requests.forEach(async (req) => {
          const request = {
            id: uuidv4(),
            name: req.name,
            method: req.method ?? "GET",
            uri_component: req.request ? req.request.url.raw : "",
            responses: [],
            collection_id: collection_copy.id,
            folder_id: folder.id,
          };

          await useFetch("http://127.0.0.1:8000/admin/request.json", {
            method: "POST",

            headers: { "Content-type": "application/json" },
            body: JSON.stringify(request),
          });
          let responses;

          if ((responses = req.response)) {
            responses.forEach(async (res) => {
              const response = {
                id: uuidv4(),
                name: res.name,
                method: res.originalRequest.method ?? "",
                uri_component: res.originalRequest.url
                  ? res.originalRequest.url.raw
                  : "",
                request_id: request.id,
                body: res.body,
              };
              await useFetch("http://127.0.0.1:8000/admin/response.json", {
                method: "POST",

                headers: { "Content-type": "application/json" },
                body: JSON.stringify(response),
              });
              request.responses.push(response);
            });
          }
          folder.requests.push(request);
          console.log(folder);
        });
        collection_copy.folders.push(folder);
      } else {
        const request = {
          id: uuidv4(),
          name: item.name,
          method: item.request.method ?? "GET",
          uri_component: item.request ? item.request.url.raw : "",
          responses: [],
          collection_id: collection_copy.id,
          folder_id: null,
        };
        console.log(item);
        console.log(request);
        await useFetch("http://127.0.0.1:8000/admin/request.json", {
          method: "POST",

          headers: { "Content-type": "application/json" },
          body: JSON.stringify(request),
        });
        if (item.response) {
          item.response.forEach((res) => {
            const response = {
              id: uuidv4(),
              name: res.name,
              method: res.originalRequest.method,
              uri_component: res.originalRequest.url.raw,
              request_id: request.id,
              body: res.body,
            };
            request.responses.push(response);
          });
        }
        collection_copy.requests.push(request);
      }
    });
  }
  props.collections?.push(collection_copy);
  uploadRef.value?.clear();
  return notification.success({
    content: "Your collection exported successfuly!",
    duration: 5000,
    closable: false,
  });
};
</script>
<style scoped>
.no-tailwind {
  margin-left: 20px;
  background-color: #ff6c37 !important;
}
.no-tailwind:hover {
  opacity: 0.8;
}
</style>
