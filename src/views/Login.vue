<template>
  <div class="bg"></div>

  <div class="login">
    <div class="floatCard">
      <div direction="vertical" class="loginForm">
        <h1 style="user-select: none; color: #2f8bcc">No Sign Up</h1>
        <Login/>
      </div>
    </div>

  </div>
</template>

<script>
import Login from '../components/Login.vue';
import {onMounted} from "vue";
import ParallaxTiltEffect from "../utils/ParrallaxTiltEffect.js";

export default {
  components: {
    Login
  },
  setup() {
    onMounted(() => {
      new ParallaxTiltEffect('.loginForm');
    });
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
}

.floatCard {
  margin: 2rem;
  width: 500px;
  height: 400px;
  perspective: 1000px;
  transform-style: preserve-3d;
  transform-origin: 0 0;
}


.loginForm {
  --X: 0;
  --Y: 0;
  margin: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  transform-style: preserve-3d;
  transform: rotateX(var(--X)) rotateY(var(--Y));
  transition: all 0.1s ease;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .2);

  flex-direction: column;
  height: 350px;
  width: 450px;
  background: linear-gradient(to bottom right, #fafafa4c, #ffffff99, #fafafa6c);
  backdrop-filter: blur(30px);
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
