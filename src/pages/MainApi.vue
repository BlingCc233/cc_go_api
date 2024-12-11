<template>
  <t-layout class="main">
    <t-aside>
      <ApiSideBar class="sidebar"/>
    </t-aside>

    <t-content>
      <router-view></router-view>
    </t-content>
  </t-layout>
</template>

<script>
import ApiSideBar from '../components/ApiSideBar.vue';

export default {
  components: {
    ApiSideBar
  }
}
</script>

<style>
.main {
  height: 100vh;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 3;
}

</style>
