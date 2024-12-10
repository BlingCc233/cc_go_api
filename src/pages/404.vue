<template>
  <div class="not-found">
    <h1>404 - 页面未找到</h1>
    <p>抱歉，您访问的页面不存在。您可以尝试以下操作：</p>
    <t-link theme="primary" @click="goBack" underline> 返回上一级 </t-link>
  </div>
</template>

<script setup>
// 引入必要的库
import { useRouter } from 'vue-router';

const router = useRouter();

const goBack = () => {
  window.history.back();
};
</script>

<style scoped>
* {
  font-family: ChiHei;
}

.not-found {
  text-align: center;
  padding: 50px;
  background-color: #f7f7f7;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.not-found h1 {
  font-size: 3em;
  color: #333;
}

.not-found p {
  font-size: 1.2em;
  color: #666;
  padding-bottom: 20px;
}

.not-found-image {
  max-width: 300px;
  margin-top: 20px;
}
</style>
