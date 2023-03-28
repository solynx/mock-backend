<template>
  <div class="sidebar w-72 h-screen bg-neutral-300">
    <nav class="sidebar__group">
      <ul class="sidebar__group__list flex py-5">
        <li class="sidebar__group_list__item sidebar__group_list__item--create">
          <n-button size="small" tertiary @click="newCollection">
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
      <NScrollbar style="max-height: 85vh">
        <ul class="sidebar__group__collection">
          <n-collapse :arrow="false" :accordion="true">
            <n-collapse-item
              v-for="(collection, index) in props.collections"
              :title="`${collection.name}`"
              :name="collection.id"
              :key="collection.id"
            >
              <template #header>
                <span @click="Alo()" v-if="editingIndex !== collection.id">
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
                  @click="exportModal"
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
                  <div><b>Collection Name</b></div>
                  <div>will be exported as a JSON file. Export as:</div>
                  <n-radio-group v-model:value="json_version" name="radiogroup">
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
                        @click="exportModal"
                      >
                        Cancel
                      </n-button>
                      <n-button strong class="no-tailwind" color="#ff6c37">
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
                    <n-button @click="toggleEdit1(collection)"> Edit </n-button>
                    <template #footer>
                      <n-button @click="deleteCollection(collection, index)">
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
                        <n-button @click="toggleEdit1(folder)"> Edit </n-button>
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
                    <n-collapse
                      v-for="(request, index) in folder.requests"
                      @click="
                        getItemSelected(
                          [
                            { id: collection.id, name: collection.name },
                            { id: folder.id, name: folder.name },
                          ],
                          request
                        )
                      "
                    >
                      <n-collapse-item
                        :title="`${request.name}`"
                        :name="request.id"
                      >
                        <template #header>
                          <span v-if="editingIndex !== request.id">
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
                                  @click="deleteRequest(folder, request, index)"
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
                            <span v-if="editingIndex !== res.id">
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
                    <span v-if="editingIndex !== request.id">
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
                            @click="deleteRequest(collection, request, index)"
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
                      <span v-if="editingIndex !== res.id">
                        {{ res.name }}</span
                      >
                      <input
                        v-else
                        type="text"
                        v-model="name"
                        @blur="handleBlur"
                        @keyup.enter="saveName(request, res, name, 'RESPONSE')"
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
} from "naive-ui";
import {
  AddOutline as AddIcon,
  FilterOutline as FilterIcon,
  EllipsisHorizontalOutline as MoreIcon,
} from "@vicons/ionicons5";
import { v4 as uuidv4 } from "uuid";
const props = defineProps({
  collections: Array,
});
const emit = defineEmits(["item-selected"]);
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
  console.log(status1);
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
  emit("item-selected", bread_cum, item);
};
const Alo = () => {
  console.log("Da");
};

const exportModal = () => {
  mode.value = !mode.value;
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
