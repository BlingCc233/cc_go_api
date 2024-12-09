<template>
  <t-layout>
    <t-content>
      <div class="bg">
        <div class="home-container">

          <!-- 介绍语 -->
          <div class="intro">
            <h1>欢迎来到我们的应用</h1>
          </div>
          <!-- 轮播图 -->
          <div class="slide-container">
            <t-swiper
                :interval="1000"
                theme="light"
                animation="fade"
                trigger="hover"
                class="custom-swiper"
                :navigation="{placement:'inside',type:'bars',showSlideBtn:'never'}"
                autoplay
                loop
                :stopOnHover="false"
            >
              <t-swiper-item v-for="(slide, index) in slides" :key="index">
                <img :src="slide.image" alt="slide.title"/>
              </t-swiper-item>
            </t-swiper>
          </div>


          <!-- 音乐播放器 -->
          <aplayer :music="{
          title: 'Song Title',
          artist: 'Artist Name',
          src: 'path/to/music.mp3',
          pic: 'path/to/cover.jpg'
        }"/>

          <!--        &lt;!&ndash; 可交互的桌宠 &ndash;&gt;-->
          <!--        <div class="desktop-pet" @click="movePet">-->
          <!--          <img src="path/to/pet.png" alt="Desktop Pet" ref="pet" />-->
          <!--        </div>-->

          <!-- 其他按钮 -->
          <t-button theme="primary" variant="base" @click="handleClick">点击我</t-button>
          <t-button theme="primary" variant="plain" to="/about">了解更多</t-button>
        </div>
      </div>
    </t-content>
  </t-layout>
</template>

<script setup>
import {onMounted} from 'vue';

const slides = ref([]);

onMounted(() => {
  for (let i = 1; i <= 18; i++) {
    if (i < 10) {
      slides.value.push({image: `src/assets/slide/output_000${i}.jpg`, title: `Slide ${i}`});
    } else if (i < 100) {
      slides.value.push({image: `src/assets/slide/output_00${i}.jpg`, title: `Slide ${i}`});
    } else if (i < 1000) {
      slides.value.push({image: `src/assets/slide/output_0${i}.jpg`, title: `Slide ${i}`});
    } else {
      slides.value.push({image: `src/assets/slide/output_${i}.jpg`, title: `Slide ${i}`});
    }
  }
});

import {ref} from 'vue';


const handleClick = () => {
  alert('按钮被点击了！');
};

const pet = ref(null);
let petPosition = {x: 0, y: 0};

const movePet = () => {
  petPosition.x += 50; // 桌宠移动的距离
  pet.value.style.transform = `translate(${petPosition.x}px, ${petPosition.y}px)`;
};
</script>


<style scoped>
* {
  border-radius: 20px;

}

.bg {
  width: 100%;
  height: 100vh;
  position: relative;
  background-color: #fafafa;
  /*background-image: url("../assets/slide/output_0001.jpg");*/
  background-repeat: no-repeat;
  background-size: cover;
}

.home-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100vh;
}

.intro {
  width: 80%;
  margin-bottom: 20px;
  padding-bottom: 20px;
  background-color: #5c5c5c;
  color: #fafafa;
}

.slide-container {
  display: flex;
  justify-content: center;
  width: 80%;
  height: 40vh;
  background-color: #5c5c5c;
  backdrop-filter: blur(300px);
  -webkit-backdrop-filter: blur(300px);
}

.t-swiper-item img {
  width: 100%;
  height: auto;
}

.button {
  margin-top: 20px;
  cursor: pointer;
}

.desktop-pet {
  position: absolute;
  bottom: 20px;
  right: 20px;
  cursor: pointer;
}

.desktop-pet img {
  width: 100px;
  height: 100px;
}

.custom-swiper {
  width: 70%; /* 设置轮播图的宽度 */
  height: 40vh; /* 设置轮播图的高度 */
}

.custom-swiper img {
  width: 100%;
  height: 30vh;
  object-fit: contain; /* 确保图片覆盖整个容器 */
  margin-top: 5vh;
}
</style>
