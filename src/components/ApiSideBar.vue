<template>
  <t-menu theme="dark" :collapsed="isCollapsed" :value="activeItem" @change="changeHandler">
    <template #logo>
      <img height="28" src="../assets/ico.png" alt="logo"/>
      <span style="font-size: 1.7rem; font-family: Log; user-select: none; padding-left: 20px">API</span>
    </template>
    <t-menu-item v-for="(item, index) in menuItems" :key="index" :value="`item${index + 1}`"
                 @click="selectItem(item.text)">
      <template #icon>
        <icon :name="item.icon"/>
      </template>
      {{ item.text }}
    </t-menu-item>
    <template #operations>
      <t-button class="t-demo-collapse-btn" variant="text" shape="square" @click="changeCollapsed">
        <template #icon>
          <t-icon name="view-list"/>
        </template>
      </t-button>
    </template>
  </t-menu>
</template>

<script lang="ts" setup>
import {ref, onMounted, watch} from 'vue';
import router from "../router/index.js";
import {useRoute} from "vue-router";
import {Icon} from 'tdesign-icons-vue-next';
import {defineEmits} from 'vue';

const route = useRoute();
const activeItem = ref<string>("");

// 定义菜单项和对应的路由
const menuItems = [
  {text: "首页", path: "/api/home", icon: "app"},
  {text: "浅草寺抽签", path: "/api/qcs", icon: "sticky-note"},
  {text: "喜报", path: "/api/xb", icon: "visual-recognition"},
  {text: "yes/no", path: "/api/yesno", icon: "chat-bubble-help"},
  {text: "QQ机器人", path: "/api/qqBot", icon: "accessibility"},
  {text: "Phi成绩", path: "/api/phib19", icon: "collage"},
  {text: "二维码生成", path: "/api/t2qr", icon: "qrcode"}
];

// 更新激活的菜单项
const updateActiveItem = () => {
  const foundItem = menuItems.find(item => item.path === route.path);
  if (foundItem) {
    activeItem.value = `item${menuItems.indexOf(foundItem) + 1}`;
  } else {
    activeItem.value = "";
  }
};

onMounted(() => {
  updateActiveItem();
});

watch(
    () => route.path,
    () => {
      updateActiveItem();
    }
);

const changeHandler = (active: string) => {
  const index = parseInt(active.slice(4)) - 1; // 提取数字部分并转换为索引
  if (index >= 0 && index < menuItems.length) {
    router.push(menuItems[index].path);
  }
};

const isCollapsed = ref<boolean>(false);
const changeCollapsed = () => {
  isCollapsed.value = !isCollapsed.value;
};

// 定义自定义事件
const emit = defineEmits(['title-selected']);

const selectItem = (title: string) => {
  emit('title-selected', title);
};
</script>

<style scoped>
* {
  outline: none;
}

.t-demo-collapse-btn {
  color: #fff;
}

.t-demo-collapse-btn:hover {
  background-color: #4b4b4b;
  border-color: transparent;
  --ripple-color: #383838;
}
</style>
