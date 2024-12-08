<template>
  <t-menu theme="light" width="250px" :value="activeItem" @change="changeHandler">
    <template #logo>
      <img height="28" src="https://tdesign.gtimg.com/site/baseLogo-light.png" alt="logo"/>
    </template>
    <t-menu-item value="item1">main</t-menu-item>
    <t-menu-item value="item2">count</t-menu-item>
    <t-menu-item value="item3">tt1</t-menu-item>
    <t-menu-item value="item4">tt2</t-menu-item>
  </t-menu>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import router from "../router/index.js";
import { useRoute } from "vue-router";

const route = useRoute();
const activeItem = ref<string>("");

const updateActiveItem = () => {
  switch (route.path) {
    case "/home":
      activeItem.value = "item1";
      break;
    case "/dbs":
      activeItem.value = "item2";
      break;
    case "/logs":
      activeItem.value = "item3";
      break;
    case "/settings":
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
      router.push("/dbs");
      break;
    case "item3":
      router.push("/logs");
      break;
    case "item4":
      router.push("/settings");
      break;
  }
};
</script>
