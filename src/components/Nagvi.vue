<template>
  <t-menu theme="light" width="232px" :value="activeItem" @change="changeHandler">
    <template #logo>
      <img height="28" src="../assets/ico.png" alt="logo"/>
      <span style="font-size: 1.7rem; font-family: Log; user-select: none">BlingCc</span>
    </template>
    <t-menu-item v-for="(item, index) in menuItems" :key="index" :value="`item${index + 1}`">
      <template #icon>
        <icon :name="item.icon"/>
      </template>
      {{ item.text }}
    </t-menu-item>
    <template #operations>
      <t-button variant="text" shape="square" @click="clearCookie">
        <template #icon><t-icon name="rollback" /></template>
      </t-button>
    </template>
  </t-menu>
</template>

<script lang="ts" setup>
import {ref, onMounted, watch} from "vue";
import router from "../router/index.js";
import {useRoute} from "vue-router";
import {Icon} from 'tdesign-icons-vue-next';
import Cookies from "js-cookie";

const route = useRoute();
const activeItem = ref<string>("");

// 定义菜单项和对应的路由
const menuItems = [
  {text: "首页", path: "/home", icon: "app"},
  {text: "心愿单", path: "/wishlist", icon: "heart"},
  {text: "tt1", path: "/test", icon: "logo-windows"},
  {text: "tt2", path: "/dbs", icon: "logo-windows"}
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

const clearCookie = () => {
  Cookies.remove('token');
  Cookies.remove('user');
  window.location.href = '/';
}
</script>

<style scoped>
*{
  outline: none;
}


</style>
