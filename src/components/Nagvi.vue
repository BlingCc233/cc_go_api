
<template>
  <t-menu theme="light"  default-value="item1" width="250px" @change="changeHandler">
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
import router from "../router/index.js";

const changeHandler = (active) => {

  console.log('change', active);
  if (active === 'item1') {
    router.push('/home');
  }
  if (active === 'item2') {
    router.push('/dbs');
  }
  if (active === 'item3') {
    router.push('/logs');
  }
  if (active === 'item4') {
    router.push('/settings');
  }
};
</script>
