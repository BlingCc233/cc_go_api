<template>
  <t-menu theme="light" width="232px" :value="activeItem" @change="changeHandler">
    <template #logo>
      <img height="28" src="../assets/ico.png" alt="logo"/>
      <span style=" font-size: 1.7rem; font-family: Log; user-select: none">BlingCc</span>
    </template>
    <t-menu-item value="item1">首页</t-menu-item>
    <t-menu-item value="item2">心愿单</t-menu-item>
    <t-menu-item value="item3">tt1</t-menu-item>
    <t-menu-item value="item4">tt2</t-menu-item>
  </t-menu>
</template>

<script lang="ts" setup>
import {ref, onMounted, watch} from "vue";
import router from "../router/index.js";
import {useRoute} from "vue-router";

const route = useRoute();
const activeItem = ref<string>("");

const updateActiveItem = () => {
  switch (route.path) {
    case "/home":
      activeItem.value = "item1";
      break;
    case "/wishlist":
      activeItem.value = "item2";
      break;
    case "/test":
      activeItem.value = "item3";
      break;
    case "/dbs":
      activeItem.value = "item4";
      break;
    default:
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
  console.log("change", active);
  switch (active) {
    case "item1":
      router.push("/home");
      break;
    case "item2":
      router.push("/wishlist");
      break;
    case "item3":
      router.push("/test");
      break;
    case "item4":
      router.push("/dbs");
      break;
  }
};
</script>
