<template>
  <div class="api-compo">
    <t-row align="middle" justify="space-between">
      <t-col :span="1">

      </t-col>
      <t-col :span="3" style="height: 100vh" class="leftDes">
        <div class="topTitle">
          <t-typography-title level="h2">{{ apiData.subtitle }}</t-typography-title>
          <t-typography-title level="h4">{{ egweb }}</t-typography-title>
          <t-typography-title level="h4">{{ egport }}</t-typography-title>
          <t-typography-title level="h2"> {{ apiData.title }}</t-typography-title>
        </div>

        <div>
          <t-typography-paragraph :ellipsis="{  row: 1, expandable: true,  collapsible: true,}">
            {{ apiData.description }}
          </t-typography-paragraph>

          <t-link theme="primary" underline :href="apiData.apiurl" target="_blank">
            <template #suffix-icon>
              <icon name="jump"/>
            </template>
            外部访问API
          </t-link>
          <p>请求类型:
            <t-typography-text code>{{ apiData.method }}</t-typography-text>
          </p>
          <p>示例参数:
            <t-typography-text code>{{ apiData.params }}</t-typography-text>
          </p>
        </div>

        <div>
          <t-button size="large" theme="primary" shape="round" variant="outline" @click="fetchApiData">
            获取响应
          </t-button>
        </div>
      </t-col>

      <t-col :span="1">
      </t-col>

      <t-col :span="7">
        <div class="tdesign-demo-image-viewer__base" v-if="imgResponse">
          <t-image-viewer :images="[imgResponse]">
            <template #trigger="{ open }">
              <div class="tdesign-demo-image-viewer__ui-image">
                <img alt="test" :src="imgResponse" class="tdesign-demo-image-viewer__ui-image--img"/>
                <div class="tdesign-demo-image-viewer__ui-image--hover" @click="open">
                  <span><icon name="browse" size="1.4em"/> 预览</span>
                </div>
              </div>
            </template>
          </t-image-viewer>
        </div>
        <div class="response">
      <pre v-if="response" :key="randKey">
              <t-button variant="text" shape="round" theme="warning" @click="copyToClipboard(response)" ghost>
                复制
              </t-button>
        <code class="language-json"></code>
      </pre>
        </div>

      </t-col>
    </t-row>
  </div>
</template>

<script setup>
import {ref, computed, watch, onMounted, nextTick} from 'vue';

import hljs from 'highlight.js/lib/core';
import 'highlight.js/styles/vs2015.css';
import 'highlight.js/lib/common'
import router from "../router/index.js";
import {Icon} from 'tdesign-icons-vue-next';
import {MessagePlugin} from 'tdesign-vue-next';


const props = defineProps({
  title: String
});

const egweb = "amywxd.site";
const egport = ":3051/api";

const apiDictionary = [
  {
    title: "qcsimg",
    subtitle: "浅草寺抽签:",
    description: "浅草寺抽签的图：\n" +
        "随机得到一张签纸，有对应的签运，解签。",
    apiurl: "http://" + egweb + egport + "/qcsimg",
    method: "GET",
    params: {}
  },
  {
    title: "qcsjson",
    subtitle: "浅草寺抽签:",
    description: "浅草寺抽签的json：\n" +
        "以json格式随机得到一张签纸，有对应的签运，解签。",
    apiurl: "http://" + egweb + egport + "/qcsjson",
    method: "GET",
    params: {}
  },
  {
    title: "xb",
    subtitle: "喜报:",
    description: "生成一张喜报。",
    apiurl: "http://" + egweb + egport + "/xb",
    method: "POST",
    params: {content: "示例"}
  },
  {
    title: "yesno",
    subtitle: "是 或 否:",
    description: "当你犹豫要不要干或者二选一的时候可以看看。",
    apiurl: "http://" + egweb + egport + "/yesno",
    method: "GET",
    params: {}
  },
  {
    title: "qqBot",
    subtitle: "QQ机器人:",
    description: "一个基于python的QQ机器人，请从外部访问",
    apiurl: "https://github.com/BlingCc233/ValoBot",
    method: "None",
    params: {}
  },
  {
    title: "phib19",
    subtitle: "Phigros成绩:",
    description: "上传你的session_token，返回你的b19成绩图。",
    apiurl: "http://" + egweb + egport + "/phib19",
    method: "POST",
    params: {session: "nkyjch88ydrg4js83bea9jyiw"}
  },
];

const apiData = computed(() => {
  const foundItem = apiDictionary.find(item => item.title === props.title);
  if (!foundItem) {
    router.push({name: 'NotFound', query: {fromApi: true}});
    return null;
  }
  return foundItem;
});

const response = ref(null);
const imgResponse = ref(null);
const randKey = ref(null);

// 监听title的变化，当title变化时重置response
watch(() => props.title, () => {
  response.value = null;
  imgResponse.value = null;
});


async function fetchApiData() {
  try {
    const url = new URL(apiData.value.apiurl);
    const options = {
      method: apiData.value.method,
      headers: {
        'Content-Type': 'application/json',
      },
    };

    if (apiData.value.method === 'GET') {
      // 添加 GET 请求的查询参数
      for (const [key, value] of Object.entries(apiData.value.params)) {
        url.searchParams.append(key, value);
      }
    } else {
      // 其他请求类型将参数放入请求体
      options.body = JSON.stringify(apiData.value.params);
    }

    const res = await fetch(url, options);

    if (res.headers.get('content-type').startsWith('image/')) {
      const blob = await res.blob();
      imgResponse.value = URL.createObjectURL(blob);
      return;
    }

    const data = await res.json();
    // 生成一个随机数
    randKey.value = Math.random().toString(36).substring(2, 15);
    response.value = JSON.stringify(data, null, 2);
  } catch (error) {
    console.error('Error fetching API data:', error);
    response.value = 'Error fetching data';
  }

  // 等待 Vue 完成 DOM 更新后执行高亮
  await nextTick();
  highlightCodeBlocks();

}


function highlightCodeBlocks() {
  document.querySelectorAll('pre code').forEach((block) => {
    // 清空内容后重新设置
    block.innerHTML = '';
    block.textContent = response.value; // 用纯文本填充
    hljs.highlightElement(block); // 高亮重新填充后的内容

  });
}


function copyToClipboard(text) {
  navigator.clipboard.writeText(text).then(() => {
    MessagePlugin.success('已复制到剪贴板');
  }).catch(err => {
    MessagePlugin.error('复制失败');
    console.error('无法复制文本: ', err);
  });
}
</script>

<style scoped>
.api-compo {
  height: 100vh;
  background-color: #eeeeee;
  z-index: 1;
  position: fixed;
  width: calc(100% - 232px);
  overflow-y: scroll;
  margin-left: -232px;
  padding-left: 232px;

}

.leftDes {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;

}

.topTitle {
  border-radius: 20px;
  background-color: #D5D5D5;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding-bottom: 20px;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .1);
}

.response {
  padding: 10px;
  border-radius: 20px;
  margin-top: 20px;
  width: 90%;
}

.response pre {
  padding: 10px;
  border-radius: 20px;
  background-color: #222;
  position: relative;
}

.response pre code {
  display: block;
  white-space: pre-wrap;
  overflow-x: auto;
  border-radius: 20px;
}

.tdesign-demo-image-viewer__ui-image {
  width: 100%;
  height: 100%;
  display: inline-flex;
  position: relative;
  justify-content: center;
  align-items: center;
  border-radius: var(--td-radius-small);
  overflow: hidden;
}

.tdesign-demo-image-viewer__ui-image--hover {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  left: 0;
  top: 0;
  opacity: 0;
  background-color: rgba(0, 0, 0, 0.6);
  color: var(--td-text-color-anti);
  line-height: 22px;
  transition: 0.2s;
}

.tdesign-demo-image-viewer__ui-image:hover .tdesign-demo-image-viewer__ui-image--hover {
  opacity: 1;
  cursor: pointer;
}

.tdesign-demo-image-viewer__ui-image--img {
  width: auto;
  height: auto;
  max-width: 100%;
  max-height: 100%;
  cursor: pointer;
  position: absolute;
}

.tdesign-demo-image-viewer__ui-image--footer {
  padding: 0 16px;
  height: 56px;
  width: 100%;
  line-height: 56px;
  font-size: 16px;
  position: absolute;
  bottom: 0;
  color: var(--td-text-color-anti);
  background-image: linear-gradient(0deg, rgba(0, 0, 0, 0.4) 0%, rgba(0, 0, 0, 0) 100%);
  display: flex;
  box-sizing: border-box;
}

.tdesign-demo-image-viewer__ui-image--title {
  flex: 1;
}

.tdesign-demo-popup__reference {
  margin-left: 16px;
}

.tdesign-demo-image-viewer__ui-image--icons .tdesign-demo-icon {
  cursor: pointer;
}

.tdesign-demo-image-viewer__base {
  width: 90%;
  height: 400px;
  margin: 10px;
  border: 4px solid var(--td-bg-color-secondarycontainer);
  border-radius: var(--td-radius-medium);
  margin-top: 20px;
  box-shadow: 0 0 25px 10px hsla(0, 0%, 0%, .1);
}

</style>
