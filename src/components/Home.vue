<template>
  <t-layout style="background-color: #F3F3F3">
    <t-content>
      <div class="home-container">
        <div class="mainLogo">
          <div class="paraCard">
            <img height="232px" src="../assets/ico.png" alt="logo"/>
          </div>
        </div>
        <div class="slogen">
          <div class="title-text">
            <span class="title-text-1-1">弗</span>
            <span class="title-text-1-2">太</span>
            <span class="title-text-1-3">ta</span>
            <span class="title-text-1-4">iii</span>
          </div>
        </div>
        <div class="countdown">
          🈚️蓄已经: {{ days }}天 {{ hours }}小时 {{ minutes }}分钟 {{ seconds }}秒
        </div>
      </div>
    </t-content>
    <t-footer style="text-align: end">
      <t-link theme="warning" @click="clearCookie" underline> 退出登录</t-link>
    </t-footer>
  </t-layout>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import ParallaxTiltEffect from "../utils/ParrallaxTiltEffect.js";
import Cookies from 'js-cookie';

const days = ref(0);
const hours = ref(0);
const minutes = ref(0);
const seconds = ref(0);

const targetDate = new Date('2024-06-09T00:00:00').getTime();

const updateCountdown = () => {
  const now = new Date().getTime();
  const distance = now - targetDate;


  days.value = Math.floor(distance / (1000 * 60 * 60 * 24));
  hours.value = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  minutes.value = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
  seconds.value = Math.floor((distance % (1000 * 60)) / 1000);

};

let interval;

onMounted(() => {
  new ParallaxTiltEffect('.paraCard');
  updateCountdown();
  interval = setInterval(updateCountdown, 1000);
});

const clearCookie = () => {
  Cookies.remove('token');
  Cookies.remove('user');
  window.location.href = '/';
}
</script>

<style scoped>
* {
  user-select: none;
}

.home-container {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  flex-direction: column;

}

.mainLogo {
  width: 300px;
  height: 300px;
  perspective: 1000px;
  transform-style: preserve-3d;
  transform-origin: 0 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.paraCard {
  --X: 0;
  --Y: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  transform-style: preserve-3d;
  transform: rotateX(var(--X)) rotateY(var(--Y));
  transition: all 0.1s ease;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .1);
  flex-direction: column;
  height: 300px;
  width: 300px;
  backdrop-filter: blur(30px);
  border-radius: 20px;
  background-color: #5c5c5c32;
}

.slogen {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 500px;
  height: 150px;
  background-color: #5c5c5c32;
  border-radius: 20px;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .1);
}

.title-text {
  padding-top: 1.5rem;
  text-align: center;
  font-family: ChiHei;
}

.title-text-1-1 {
  color: #474747;
  font-size: 6rem;
}

.title-text-1-2 {
  color: #474747;
  font-size: 3rem;
}

.title-text-1-3 {
  color: #DC624E;
  font-family: Log;
  font-size: 4.5rem;
}

.title-text-1-4 {
  color: #fafafa;
  font-family: Log;
  font-size: 6rem;
}


.countdown {
  text-align: center;
  background-color: #5c5c5c32;
  color: #474747;
  padding: 2rem 1.1rem 2rem 1.2rem;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .1);
  border-radius: 20px;
  margin-top: 20px;
  font-size: 2rem;
  font-family: ChiHei;
}

</style>
