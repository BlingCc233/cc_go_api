<template>
  <Nagvi class="sidebar" />
  <div class="content">
    <router-view></router-view>
  </div>
</template>

<script>
import Nagvi from '../components/Nagvi.vue';

export default {
  components: {
    Nagvi
  }
}
</script>

<style scoped>
  .sidebar {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    width: 250px;
    padding: 16px;
  }
  .content{
    margin-left: 250px;
  }
</style>
