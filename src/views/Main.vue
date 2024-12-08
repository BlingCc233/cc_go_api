<template>

  <t-layout class="main">
    <t-aside>
      <Nagvi class="sidebar"/>

    </t-aside>

    <t-content>
      <div class="content">
        <router-view>
        </router-view>
      </div>
    </t-content>

  </t-layout>

</template>

<script>
import Nagvi from '../components/Nagvi.vue';
import Home from '../components/Home.vue';

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
  width: 100vw;
}
</style>


