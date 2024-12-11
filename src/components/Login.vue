<template>
  <div style="width: 350px">
    <t-form ref="form" :data="formData" :colon="true" :label-width="0" @reset="onReset" @submit="onSubmit" style="text-align: center">
      <t-form-item name="account">
        <t-input size="large" v-model="formData.Username" clearable placeholder="请输入账户名">
          <template #prefix-icon>
            <desktop-icon/>
          </template>
        </t-input>
      </t-form-item>

      <t-form-item name="password">
        <t-input size="large" v-model="formData.password" type="password" clearable placeholder="请输入密码">
          <template #prefix-icon>
            <lock-on-icon/>
          </template>
        </t-input>
      </t-form-item>

      <button submit class="loginButton">
        登录
      </button>
    </t-form>
  </div>
</template>
<script setup>
import {onMounted, reactive} from 'vue';
import {MessagePlugin} from 'tdesign-vue-next';
import {DesktopIcon, LockOnIcon} from 'tdesign-icons-vue-next';
import router from '../router/index.js';
import Cookies from 'js-cookie';

const formData = reactive({
  Username: '',
  password: '',
});


const onReset = () => {
  MessagePlugin.success('重置成功');
};
const onSubmit = async () => {

  try {
    const response = await fetch('http://localhost:3051/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (response.ok) {
      const data = await response.json();
      MessagePlugin.success('登录成功');
      await Cookies.set('token', data.token, {expires: 7});
      await Cookies.set('user', formData.Username, {expires:7});

      router.push('/home');
      console.log(data);
    } else {
      MessagePlugin.error('登录失败');
      console.error('登录失败:', response.statusText);
    }
  } catch (error) {
    MessagePlugin.error('请求错误');
    console.error('请求错误:', error);
  }
};

</script>

<style scoped>
.loginButton {
  width: 100px;
  background-color: #2f8bcc;
  color: #fff;
  border: none;
  border-radius: 90px;
  outline: none;
  margin-top: 35px;
  font-size: 1.1rem;
  padding: 10px 0 10px 0;
  transition: all 0.5s ease;
}
.loginButton:hover {
  background-color: #2f8bcccc;
  cursor: pointer;
}

.loginButton:active {
  background-color: #2f8bcc;
}
</style>
