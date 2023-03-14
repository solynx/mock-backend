<template>
  <div class="sidebar w-64 h-screen bg-neutral-300">
    <nav class="sidebar__group">
      <ul class="sidebar__group__list flex py-5">
        <li class="sidebar__group_list__item sidebar__group_list__item--create">
          <n-icon class="sidebar__group__list--icon text-2xl">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
              <path
                d="M368.5 240H272v-96.5c0-8.8-7.2-16-16-16s-16 7.2-16 16V240h-96.5c-8.8 0-16 7.2-16 16 0 4.4 1.8 8.4 4.7 11.3 2.9 2.9 6.9 4.7 11.3 4.7H240v96.5c0 4.4 1.8 8.4 4.7 11.3 2.9 2.9 6.9 4.7 11.3 4.7 8.8 0 16-7.2 16-16V272h96.5c8.8 0 16-7.2 16-16s-7.2-16-16-16z"
              />
            </svg>
          </n-icon>
        </li>
        <li class="sidebar__group__list__item">
          <n-input-group>
            <n-input-group-label
              ><n-icon class="sidebar__group__list--icon text-2xl">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                  viewBox="0 0 24 24"
                >
                  <path
                    d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"
                    fill="currentColor"
                  ></path></svg></n-icon
            ></n-input-group-label>
            <n-input :style="{ width: '100%%' }" />
          </n-input-group>
        </li>
        <li class="sidebar__group__list__item">
          <n-icon class="sidebar__group__list--icon text-2xl">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              xmlns:xlink="http://www.w3.org/1999/xlink"
              viewBox="0 0 24 24"
            >
              <g
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <circle cx="5" cy="12" r="1"></circle>
                <circle cx="12" cy="12" r="1"></circle>
                <circle cx="19" cy="12" r="1"></circle>
              </g>
            </svg>
          </n-icon>
        </li>
      </ul>
      <ul class="sidebar__group__collection">
        <n-collapse :arrow="false" :accordion="true">
          <n-collapse-item title="Collection" name="1">
            <template #header-extra>
              <n-space>
                <NDropdown
                  trigger="click"
                  :options="treeOptions"
                  :show-arrow="true"
                >
                  <NButton
                    ><n-icon class="sidebar__group__list--icon text-2xl">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        xmlns:xlink="http://www.w3.org/1999/xlink"
                        viewBox="0 0 24 24"
                      >
                        <g
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        >
                          <circle cx="5" cy="12" r="1"></circle>
                          <circle cx="12" cy="12" r="1"></circle>
                          <circle cx="19" cy="12" r="1"></circle>
                        </g>
                      </svg> </n-icon
                  ></NButton>
                </NDropdown>
              </n-space>
            </template>
            <n-collapse>
              <n-collapse-item title="Folder 1" name="1">
                <template #header-extra>
                  <n-space>
                    <NDropdown
                      trigger="click"
                      :options="folderOptions"
                      :show-arrow="true"
                    >
                      <NButton
                        ><n-icon class="sidebar__group__list--icon text-2xl">
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            xmlns:xlink="http://www.w3.org/1999/xlink"
                            viewBox="0 0 24 24"
                          >
                            <g
                              fill="none"
                              stroke="currentColor"
                              stroke-width="2"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            >
                              <circle cx="5" cy="12" r="1"></circle>
                              <circle cx="12" cy="12" r="1"></circle>
                              <circle cx="19" cy="12" r="1"></circle>
                            </g>
                          </svg> </n-icon
                      ></NButton>
                    </NDropdown>
                  </n-space>
                </template>
                <n-collapse>
                  <n-collapse-item title="RequestABCs" name="1">
                    <template #header-extra>
                      <n-space>
                        <NDropdown
                          trigger="click"
                          :options="requestOptions"
                          :show-arrow="true"
                        >
                          <NButton
                            ><n-icon
                              class="sidebar__group__list--icon text-2xl"
                            >
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                xmlns:xlink="http://www.w3.org/1999/xlink"
                                viewBox="0 0 24 24"
                              >
                                <g
                                  fill="none"
                                  stroke="currentColor"
                                  stroke-width="2"
                                  stroke-linecap="round"
                                  stroke-linejoin="round"
                                >
                                  <circle cx="5" cy="12" r="1"></circle>
                                  <circle cx="12" cy="12" r="1"></circle>
                                  <circle cx="19" cy="12" r="1"></circle>
                                </g>
                              </svg> </n-icon
                          ></NButton>
                        </NDropdown>
                      </n-space>
                    </template>
                    <div class="flex justify-between">
                      <span class="response self-center">Respone 1</span>
                      <n-space>
                        <NDropdown
                          trigger="click"
                          :options="responseOptions"
                          :show-arrow="true"
                        >
                          <NButton class="block w-full h-full"
                            ><n-icon
                              class="sidebar__group__list--icon text-2xl block"
                            >
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                xmlns:xlink="http://www.w3.org/1999/xlink"
                                viewBox="0 0 24 24"
                              >
                                <g
                                  fill="none"
                                  stroke="currentColor"
                                  stroke-width="2"
                                  stroke-linecap="round"
                                  stroke-linejoin="round"
                                >
                                  <circle cx="5" cy="12" r="1"></circle>
                                  <circle cx="12" cy="12" r="1"></circle>
                                  <circle cx="19" cy="12" r="1"></circle>
                                </g>
                              </svg> </n-icon
                          ></NButton>
                        </NDropdown>
                      </n-space>
                    </div>
                  </n-collapse-item>
                  <div>Request1</div>
                  <n-collapse-item title="Request 2" name="2">
                    <template #header-extra>
                      <n-space>
                        <NDropdown
                          trigger="click"
                          :options="requestOptions"
                          :show-arrow="true"
                        >
                          <NButton
                            ><n-icon
                              class="sidebar__group__list--icon text-2xl"
                            >
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                xmlns:xlink="http://www.w3.org/1999/xlink"
                                viewBox="0 0 24 24"
                              >
                                <g
                                  fill="none"
                                  stroke="currentColor"
                                  stroke-width="2"
                                  stroke-linecap="round"
                                  stroke-linejoin="round"
                                >
                                  <circle cx="5" cy="12" r="1"></circle>
                                  <circle cx="12" cy="12" r="1"></circle>
                                  <circle cx="19" cy="12" r="1"></circle>
                                </g>
                              </svg> </n-icon
                          ></NButton>
                        </NDropdown>
                      </n-space>
                    </template>
                    <div>Res Json1</div>
                  </n-collapse-item>
                  <n-collapse-item title="right" name="3">
                    <div>very good</div>
                  </n-collapse-item>
                </n-collapse>
              </n-collapse-item>
              <n-collapse-item title="right" name="2">
                <div>nice</div>
              </n-collapse-item>
              <n-collapse-item title="right" name="3">
                <div>very good</div>
              </n-collapse-item>
            </n-collapse>
          </n-collapse-item>
          <n-collapse-item title="right" name="2">
            <div>nice</div>
          </n-collapse-item>
          <n-collapse-item title="right" name="3">
            <div>very good</div>
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
} from "naive-ui";
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
</script>
