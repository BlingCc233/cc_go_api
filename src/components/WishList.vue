<template>
  <div class="wishContainer">
    <t-typography-title level="h2">心愿单</t-typography-title>

    <div class="inputContent">
      <t-input clearable size="large" v-model="newWishContent" auto-width placeholder="输入新的心愿"/>
      <t-button theme="primary" size="large" type="submit" shape="round" @click="addWish(newWishContent)">确定
      </t-button>
    </div>
    <t-list stripe header="久别" footer="必重逢" class="wishContent" :scroll="{ type: 'virtual' }">
      <t-list-item v-for="wish in wishes.data" :key="wish.id">
        <t-list-item-meta :image="avatar(wish.creater)" :title="wish.content" :description="wish.creater"/>
        <template #action>
      <span>
        <input type="checkbox" :checked="wish.checked" @change="toggleCheck(wish.id)" style="margin-right: 8px"/>
        <t-button theme="primary" variant="text" @click="deleteWish(wish.id)">删除</t-button>
        <t-button theme="primary" variant="text" @click="editWish(wish)">编辑</t-button>
      </span>
        </template>
      </t-list-item>
    </t-list>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import Cookies from 'js-cookie'
import {MessagePlugin} from "tdesign-vue-next";

const wishes = ref([])
const newWishContent = ref('')

const avatar = (user) => {
  if (!user) {
    return 'src/assets/ico.png'
  }
  if (user === 'blingcc')
    return 'src/assets/ccavatar.JPG'
  if (user === 'lzhx')
    return 'src/assets/lzhxavatar.JPG'
  return `src/assets/ico.png`
}

// 检查权限
async function checkAuth() {
  // 附上Cookie的token作为header的Authorization
  const response = await fetch('http://localhost:3051/api/auth', {
    headers: {
      Authorization: `${Cookies.get('token')}`
    }
  })

  if (!response.ok) {
    throw new Error('Permission denied')
  }
}

// 获取心愿单列表
async function fetchWishes() {
  try {
    const response = await fetch('http://localhost:3051/api/get_wishes', {
      headers: {
        Authorization: `${Cookies.get('token')}`
      }
    })
    if (!response.ok) {
      throw new Error('Failed to fetch wishes')
    }
    wishes.value = await response.json()
    console.log(wishes.value.data)
  } catch (error) {
    console.error('Error fetching wishes:', error)
  }
}

// 添加心愿单
async function addWish(content) {
  if (!content.trim()) {
    MessagePlugin.warning('请输入心愿内容')
    return
  }

  try {
    const response = await fetch('http://localhost:3051/api/new_wish', {
      method: 'POST',
      headers: {
        'Authorization': `${Cookies.get('token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({'content': content, 'creater': Cookies.get('user')})
    })
    if (!response.ok) {
      throw new Error('Failed to add wish')
    }
    const newWish = await response.json()
    wishes.value.data.push(newWish) // 先更新本地数据
    // fetchWishes() // 刷新列表，这里可以选择是否调用
  } catch (error) {
    console.error('Error adding wish:', error)
  }
  newWishContent.value = '' // 清空输入框
}

// 删除心愿单
async function deleteWish(id) {
  try {
    // await checkAuth()
    const response = await fetch(`http://localhost:3051/api/delete_wish/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `${Cookies.get('token')}`
      }
    })
    if (!response.ok) {
      throw new Error('Failed to delete wish')
    }
    fetchWishes() // 刷新列表
  } catch (error) {
    console.error('Error deleting wish:', error)
  }
}


// 更新心愿单
async function updateWish(id, content, checked) {
  try {
    await checkAuth()
    const response = await fetch(`http://localhost:3051/api/update_wish/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({content, checked})
    })
    if (!response.ok) {
      throw new Error('Failed to update wish')
    }
    fetchWishes() // 刷新列表
  } catch (error) {
    console.error('Error updating wish:', error)
  }
}

// 切换心愿单的完成状态
async function toggleCheck(id) {
  const wish = wishes.value.data.find(wish => wish.id === id)
  if (wish) {
    try {
      // await checkAuth()
      const response = await fetch(`http://localhost:3051/api/wish_check/${id}`, {
        method: 'PUT',
        headers: {
          'Authorization': `${Cookies.get('token')}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({checked: !wish.checked})
      })
      if (!response.ok) {
        throw new Error('Failed to update wish status')
      }
      // 更新本地数据
      wish.checked = !wish.checked
    } catch (error) {
      console.error('Error toggling wish check:', error)
    }
  }
}


// 初始化获取心愿单列表
fetchWishes()

// 编辑心愿单
function editWish(wish) {
  const newContent = prompt('编辑心愿', wish.content)
  if (newContent !== null) {
    updateWish(wish.id, newContent, wish.checked)
  }
}
</script>

<style scoped>
.wishContainer {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
  background-color: #F3F3F3;
}
.inputContent{
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  width: 500px;

}
.wishContent {
  height: 70vh;
  width: 70%;
  border-radius: 20px;
}
</style>
