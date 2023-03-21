<template>
  <div class="sidebar w-72 h-screen bg-neutral-300">
    <nav class="sidebar__group">
      <ul class="sidebar__group__list flex py-5">
        <li class="sidebar__group_list__item sidebar__group_list__item--create">
          <n-button size="small" tertiary @click="newCollection()">
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
      <ul class="sidebar__group__collection">
        <n-collapse
          :arrow="false"
          :accordion="true"
          v-for="(collection, index) in collections12"
        >
          <n-collapse-item
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
                @keyup.enter="saveName(null, collection, name, 'COLLECTION')"
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
            <n-collapse v-for="(folder, index) in collection.folders">
              <n-collapse-item :title="`${folder.name}`" :name="folder.name">
                <template #header>
                  <span v-if="editingIndex !== folder.id">
                    {{ folder.name }}</span
                  >
                  <input
                    v-else
                    type="text"
                    v-model="name"
                    @blur="handleBlur"
                    @keyup.enter="saveName(collection, folder, name, 'FOLDER')"
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
                  <n-collapse v-for="(request, index) in folder.requests">
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
                          @keyup.enter="
                            saveName(folder, request, name, 'REQUEST')
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
                              <n-button @click="addResponse(request)">
                                Add Response
                              </n-button>
                            </template>
                            <n-button @click="toggleEdit1(request)">
                              Edit
                            </n-button>
                            <template #footer>
                              <n-button @click="deleteRequest(folder, index)">
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
                </n-collapse>
              </n-collapse-item>
            </n-collapse>
            <!-- REQUEST IN COLLECTION-->
            <n-collapse v-for="(request, index) in collection.requests">
              <n-collapse-item :title="`${request.name}`" :name="request.id">
                <template #header>
                  <span v-if="editingIndex !== request.id">
                    {{ request.name }}</span
                  >
                  <input
                    v-else
                    type="text"
                    v-model="name"
                    @blur="handleBlur"
                    @keyup.enter="
                      saveName(collection, request, name, 'REQUEST')
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
                        <n-button @click="addResponse(request)">
                          Add Response
                        </n-button>
                      </template>

                      <n-button @click="toggleEdit1(request)"> Edit </n-button>
                      <template #footer>
                        <n-button @click="deleteRequest(collection, index)">
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
                    <span v-if="editingIndex !== res.id"> {{ res.name }}</span>
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
                          <n-button @click="toggleEdit1(res)"> Edit </n-button>
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
} from "naive-ui";
import {
  AddOutline as AddIcon,
  FilterOutline as FilterIcon,
  EllipsisHorizontalOutline as MoreIcon,
} from "@vicons/ionicons5";
import { v4 as uuidv4 } from "uuid";
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

var collections1 = [
  {
    name: "New Collection",
    id: "col_abcxyz111",
    folders: [
      {
        name: "folder1",
        id: "aabbcc12",
        requests: [],
      },
      {
        name: "folder2",
        id: "fd_aabbccaz",
        requests: [
          {
            name: "Request1",
            method: "Post",
            id: "req_ndnd",
            uri: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                id: "res_aksndasndf",
              },
              {
                name: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                id: "res_aksdkas2",
              },
            ],
          },
          {
            name: "Request123",
            method: "Post",
            id: "req_ndnd",
            uri: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                id: "res_aksndasnd12",
              },
              {
                name: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                id: "res_aksdkas1",
              },
            ],
          },
        ],
      },
    ],
    requests: [
      {
        name: "Request A",
        id: "clt_req_1",
        method: "Patch",
        uri: "facebook.com",
        responses: [
          {
            name: "Response 1",
            id: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
          {
            name: "Response 1",
            id: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
        ],
      },
    ],
  },
  {
    name: "New Collection",
    id: "col_abcxyz45",
    folders: [
      {
        name: "folder1",
        id: "aabbcc12354",
      },
      {
        name: "folder2",
        id: "fd_aabbccaz234",
        requests: [
          {
            name: "Request1",
            method: "Post",
            id: "req_ndnd234",
            uri: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                id: "res_aksndasndf1",
              },
              {
                name: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                id: "res_aksdkas21",
              },
            ],
          },
          {
            name: "Request123",
            method: "Post",
            id: "req_ndnd12",
            uri: "google.com",
            responses: [
              {
                name: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                id: "res_aksndasnd123",
              },
              {
                name: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                id: "res_aksdkas12",
              },
            ],
          },
        ],
      },
    ],
  },
  {
    name: "New Collection2",
    id: "abcxyz1",
  },
];
//get all collection
// const { data: collections12 } = await useFetch(
//   "http://127.0.0.1:8000/admin/collection.json",
//   {}
// );
// const collections123 = JSON.parse(JSON.stringify(collections12.value));

// const collections = ref(collections123.data);

const folderA = [
  {
    label: "Edit",
    key: "key1",
  },
  {
    label: "Delete",
    key: "key2",
  },
  {
    label: "Add Request",
    key: "key23",
  },
];
const requestOptions = [
  {
    label: "Edit",
    key: "key1",
  },
  {
    label: "Delete",
    key: "key2",
  },
  {
    label: "Add Response",
    key: "key3",
  },
];
const responseOptions = [
  {
    label: "Edit",
    key: "key1",
  },
  {
    label: "Delete",
    key: "key2",
  },
];
const NearFilterOptions = [
  {
    label: "Open Trash",
    key: "key1",
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
const collections12 = ref(collections1);
const mode = ref(true);

const newCollection = () => {
  const uuid = uuidv4();
  var collection = {
    name: "New Collection",
    id: uuid,
    folders: [],
    requests: [],
  };
  collections12.value.push(collection);
  // const { data: status } = await useFetch(
  //   "http://127.0.0.1:8000/admin/collection.json",
  //   {
  //     method: "POST",
  //     headers: { "Content-type": "application/json" },
  //     body: JSON.stringify(collection),
  //   }
  // );
  // const status1 = JSON.parse(JSON.stringify(status.value));
  // console.log(status1);
  // if (status1.status) {
  //   collections.value.push(collection);
  // } else {
  //   return alert("Created collection failed");
  // }
};
const toggleEdit = true;
const optionName = ref(null);
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
const Alo = () => {
  console.log("A");
};
const addRequest = (collection: object, folder?: object) => {
  const uuid = uuidv4();
  if (folder) {
    const request = {
      id: uuid,
      name: "New Request",
      folder_id: folder.id,
      collection_id: collection.id,
      url_component: "",
      method: "GET",
      responses: [],
    };
    folder.requests.push(request);
    return;
  }
  const request = {
    id: uuid,
    name: "New Request",
    folder_id: null,
    collection_id: collection.id,
    url_component: "",
    method: "GET",
    responses: [],
  };
  collection.requests.push(request);
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
const deleteFolder = (collection: object, folder: object, index: number) => {
  collection.folders.splice(index, 1);
};
const newFolder = (collection: object) => {
  const uuid = uuidv4();
  const folder = {
    name: "New Folder",
    id: uuid,
    requests: [],
  };
  collection.folders.push(folder);
};
const deleteCollection = (collection: object, index: number) => {
  collections12.value.splice(index, 1);
};
const deleteRequest = (parent_el: object, index: number) => {
  parent_el.requests.splice(index, 1);
};
const addResponse = (request: object) => {
  const uuid = uuidv4();
  const response = {
    id: uuid,
    name: "New response",
    method: "GET",
    uri_component: null,
    request_id: request.id,
    file_link: null,
  };

  request.responses.push(response);
};
const deleteResponse = (request: object, response: object, index: number) => {
  request.responses.splice(index, 1);
};
const editResponse = (response: object, new_name: string) => {
  response.name = new_name;
};
</script>
