<template>
  <div>
    <h1>心愿单</h1>
    <input v-model="newWishContent" placeholder="输入新的心愿" />
    <button @click="addWish(newWishContent)">添加心愿</button>
    <ul>
      <li v-for="wish in wishes.data" :key="wish.id">
        <input type="checkbox" :checked="wish.checked" @change="toggleCheck(wish.id)" />
        <span :style="{ textDecoration: wish.checked ? 'line-through' : 'none' }">{{ wish.content }}</span>
        <button @click="deleteWish(wish.id)">删除</button>
        <button @click="editWish(wish)">编辑</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Cookies from 'js-cookie'

const wishes = ref([])
const newWishContent = ref('')

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
    alert('请输入心愿内容')
    return
  }

  try {
    const response = await fetch('http://localhost:3051/api/new_wish', {
      method: 'POST',
      headers: {
        'Authorization': `${Cookies.get('token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 'content': content, 'creater': Cookies.get('user') })
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
    await checkAuth()
    const response = await fetch(`http://localhost:3051/api/delete_wish/${id}`, {
      method: 'DELETE'
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
      body: JSON.stringify({ content, checked })
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
function toggleCheck(id) {
  const wish = wishes.value.find(wish => wish.id === id)
  updateWish(id, wish.content, !wish.checked)
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
/* 添加一些样式 */
</style>
