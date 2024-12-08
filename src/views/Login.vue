<template>
  <div class="bg"></div>

  <div class="login">
    <div  direction="vertical" class="loginForm">
      <h1 style="user-select: none; color: #F7DAB1">No Sign Up</h1>
      <Login/>
    </div>
  </div>
</template>

<script>
import Login from '../components/Login.vue';

export default {
  components: {
    Login
  }
}
</script>

<style scoped>
.login {
  display: flex;
  place-items: center;
  flex-direction: column;
  justify-content: center;
  height: 100vh;
  position: relative; /* 添加相对定位 */
  z-index: 1; /* 确保登录表单在背景之上 */
  backdrop-filter: blur(50px); /* 添加毛玻璃效果 */
}

.loginForm{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 30vh;
  width: 450px;
  background-color: rgba(255, 255, 255, 0.3);
  padding-bottom: 20px;
  border-radius: 20px;
}

.bg {
  background-image: url("../assets/login_bg.JPG");
  background-size: contain;
  position: fixed; /* 固定背景 */
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0; /* 确保背景在登录表单之下 */
}
</style>
