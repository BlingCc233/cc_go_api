<template>

  <t-layout class="main">
    <t-aside>
      <Nagvi class="sidebar"/>

    </t-aside>

    <t-content>
      <router-view>
      </router-view>
    </t-content>

  </t-layout>

</template>

<script>
import Nagvi from '../components/ApiSideBar.vue';
import Home from '../components/Home.vue'

export default {
  components: {
    Nagvi,
    Home
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
}
</style>


