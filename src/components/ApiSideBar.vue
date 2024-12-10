<template>
  <t-menu theme="light" width="232px" :value="activeItem" @change="changeHandler">
    <template #logo>
      <img height="28" src="../assets/ico.png" alt="logo"/>
      <span style="font-size: 1.7rem; font-family: Log; user-select: none">API</span>
    </template>
    <t-menu-item v-for="(item, index) in menuItems" :key="index" :value="`item${index + 1}`">
      {{ item.text }}
    </t-menu-item>
  </t-menu>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import router from "../router/index.js";
import { useRoute } from "vue-router";

const route = useRoute();
const activeItem = ref<string>("");

// 定义菜单项和对应的路由
const menuItems = [
  { text: "首页", path: "/api/home" },
  { text: "心愿单", path: "/api/wishlist" },
  { text: "tt1", path: "/api/test" },
  { text: "tt2", path: "/api/dbs" }
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
</script>

<style>
.main {
  height: 100vh;
}
.sidebar {
  position: fixed;
  top: 0;
  left: 0;
}
</style>
