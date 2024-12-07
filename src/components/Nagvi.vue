<template>
  <div class="box">
    <t-menu theme="light" default-value="item1" style="margin-right: 40px" @change="changeHandler">
      <template #logo>
        <img height="28" src="https://tdesign.gtimg.com/site/baseLogo-light.png" alt="logo" />
      </template>
      <t-menu-item value="item1"> 仪表盘 </t-menu-item>
      <t-menu-item value="item2"> 资源列表 </t-menu-item>
      <t-menu-item value="item3"> 根目录 </t-menu-item>
      <t-menu-item value="item4" :disabled="true"> 调度平台 </t-menu-item>
      <t-menu-item value="item5"> 精准监控 </t-menu-item>
      <t-menu-item value="item6"> 消息区 </t-menu-item>
      <t-menu-item value="item7"> 个人中心 </t-menu-item>
      <t-menu-item value="item8"> 视频区 </t-menu-item>
      <t-menu-item value="item9"> 资源编辑 </t-menu-item>
    </t-menu>


  </div>
</template>

<script setup>
import router from "../router/index.js";
const changeHandler = (active) => {

  console.log('change', active);
  if (active === 'item1'){
    router.push('/dbs');
  }
};
</script>
