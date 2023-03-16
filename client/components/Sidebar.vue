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
        <n-collapse :arrow="false" :accordion="true" v-for="collection in col">
          <n-collapse-item
            :title="`${collection.label}`"
            :name="collection.key"
          >
            <!-- FOLDERS IN COLLECTION-->
            <n-collapse v-for="folder in collection.folders">
              <n-collapse-item :title="`${folder.label}`" :name="folder.label">
                <template #header-extra>
                  <n-space>
                    <NDropdown
                      trigger="click"
                      :options="folderOptions"
                      :show-arrow="true"
                    >
                      <n-button size="medium" tertiary>
                        <n-icon>
                          <more-icon></more-icon>
                        </n-icon>
                      </n-button>
                    </NDropdown>
                  </n-space>
                </template>
                <n-collapse>
                  <!-- Request In Folder-->
                  <n-collapse v-for="request in folder.requests">
                    <n-collapse-item
                      :title="`${request.label}`"
                      :name="request.key"
                    >
                      <template #header-extra>
                        <n-space>
                          <NDropdown
                            trigger="click"
                            :options="requestOptions"
                            :show-arrow="true"
                          >
                            <n-button>
                              <n-icon>
                                <more-icon></more-icon>
                              </n-icon>
                            </n-button>
                          </NDropdown>
                        </n-space>
                      </template>
                      <!-- RESPONSE IN REQUEST OF folders-->

                      <n-collapse-item
                        v-for="res in request.response"
                        :title="`${res.label}`"
                        :name="res.key"
                      >
                        <template #header-extra>
                          <n-space>
                            <NDropdown
                              trigger="click"
                              :options="responseOptions"
                              :show-arrow="true"
                            >
                              <n-button>
                                <n-icon>
                                  <more-icon></more-icon>
                                </n-icon>
                              </n-button>
                            </NDropdown>
                          </n-space>
                        </template>
                      </n-collapse-item>
                    </n-collapse-item>
                  </n-collapse>
                </n-collapse>
              </n-collapse-item>
            </n-collapse>
            <!-- REQUEST IN COLLECTION-->
            <n-collapse v-for="request in collection.requests">
              <n-collapse-item :title="`${request.label}`" :name="request.key">
                <template #header-extra>
                  <n-space>
                    <NDropdown
                      trigger="click"
                      :options="requestOptions"
                      :show-arrow="true"
                    >
                      <n-button>
                        <n-icon>
                          <more-icon></more-icon>
                        </n-icon>
                      </n-button>
                    </NDropdown>
                  </n-space>
                </template>
                <!-- RESPONSE IN REQUEST OF COLLECTIONs-->

                <n-collapse-item
                  v-for="res in request.response"
                  :title="`${res.label}`"
                  :name="res.key"
                >
                  <template #header-extra>
                    <n-space>
                      <NDropdown
                        trigger="click"
                        :options="responseOptions"
                        :show-arrow="true"
                      >
                        <n-button>
                          <n-icon>
                            <more-icon></more-icon>
                          </n-icon>
                        </n-button>
                      </NDropdown>
                    </n-space>
                  </template>
                </n-collapse-item>
              </n-collapse-item>
            </n-collapse>
          </n-collapse-item>
          <template #header-extra>
            <n-space>
              <NDropdown
                trigger="click"
                :options="treeOptions"
                :show-arrow="true"
              >
                <n-button tertiary size="small">
                  <n-icon size="24">
                    <more-icon />
                  </n-icon>
                </n-button>
              </NDropdown>
            </n-space>
          </template>
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
} from "naive-ui";
import {
  AddOutline as AddIcon,
  FilterOutline as FilterIcon,
  EllipsisHorizontalOutline as MoreIcon,
} from "@vicons/ionicons5";
const treeOptions = [
  {
    label: "Add Folder",
    key: "key1",
  },
  {
    label: "Edit",
    key: "key2",
  },
  {
    label: "Delete",
    key: "key2",
  },
];

var collections = [
  {
    label: "New Collection",
    key: "col_abcxyz",
    folders: [
      {
        label: "folder1",
        key: "aabbcc12",
      },
      {
        label: "folder2",
        key: "fd_aabbccaz",
        requests: [
          {
            label: "Request1",
            method: "Post",
            key: "req_ndnd",
            uri: "google.com",
            response: [
              {
                label: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                key: "res_aksndasndf",
              },
              {
                label: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                key: "res_aksdkas2",
              },
            ],
          },
          {
            label: "Request123",
            method: "Post",
            key: "req_ndnd",
            uri: "google.com",
            response: [
              {
                label: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                key: "res_aksndasnd12",
              },
              {
                label: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                key: "res_aksdkas1",
              },
            ],
          },
        ],
      },
    ],
    requests: [
      {
        label: "Request A",
        key: "clt_req_1",
        method: "Patch",
        uri: "facebook.com",
        response: [
          {
            label: "Response 1",
            key: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
          {
            label: "Response 1",
            key: "col_req_resABC",
            method: "patch",
            json_link: "json file link in storage",
          },
        ],
      },
    ],
  },
  {
    label: "New Collection",
    key: "col_abcxyz45",
    folders: [
      {
        label: "folder1",
        key: "aabbcc12354",
      },
      {
        label: "folder2",
        key: "fd_aabbccaz234",
        requests: [
          {
            label: "Request1",
            method: "Post",
            key: "req_ndnd234",
            uri: "google.com",
            response: [
              {
                label: "Response 1",
                method: "Delete",
                json_file: "link json file in storage",
                key: "res_aksndasndf1",
              },
              {
                label: "Response 1",
                method: "Post",
                json_file: "link json file in storage",
                key: "res_aksdkas21",
              },
            ],
          },
          {
            label: "Request123",
            method: "Post",
            key: "req_ndnd12",
            uri: "google.com",
            response: [
              {
                label: "Response 1",
                method: "Patch",
                json_file: "link json file in storag11e",
                key: "res_aksndasnd123",
              },
              {
                label: "Response 13",
                method: "Post",
                json_file: "link json file in storage1",
                key: "res_aksdkas12",
              },
            ],
          },
        ],
      },
    ],
  },
  {
    label: "New Collection2",
    key: "abcxyz1",
  },
];
var col = ref(collections);
const folderOptions = [
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
    key: "key2",
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
const toggle = ref(false);
const toggleButton = (status: any) => {
  toggle.value = !status.value;
  console.log(toggle.value);
};
const newCollection = () => {
  console.log(col);

  var collection = {
    label: "New Collection",
    key: "abcs",
    folders: [],
    requests: [],
  };
  col.value.push(collection);
};
</script>
