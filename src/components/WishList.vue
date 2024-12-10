<template>
  <div class="wishContainer">
    <t-typography-title level="h2">心愿单</t-typography-title>

    <t-space direction="vertical" class="inputContent">
      <t-textarea placeholder="输入新的心愿"
                  v-model="newWishContent" autofocus autosize
      />
      <t-button theme="primary" type="submit" shape="round" @click="addWish(newWishContent)">确定
      </t-button>
    </t-space>
    <t-list stripe header="久别" footer="必重逢" class="wishContent" :scroll="{ type: 'virtual' }">
      <t-list-item v-for="wish in wishes.data" :key="wish.id">
        <t-list-item-meta :image="avatar(wish.creater)" :title="wish.content" :description="wish.creater"/>
        <template #action>
          <t-button variant="text" shape="square" style="margin-right: 15rem">
            <input type="checkbox" :checked="wish.checked" @change="toggleCheck(wish.id)"/>
          </t-button>
          <t-button variant="text" shape="square" @click="deleteWish(wish.id)">
            <delete-icon/>
          </t-button>
          <t-button variant="text" shape="square" @click="editWish(wish)">
            <edit-icon/>
          </t-button>

        </template>
      </t-list-item>
    </t-list>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import Cookies from 'js-cookie'
import {MessagePlugin} from "tdesign-vue-next";
import {DeleteIcon, EditIcon, DownloadIcon} from 'tdesign-icons-vue-next';

const wishes = ref([])
const newWishContent = ref('')
const user = Cookies.get('user')

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
    sortWishes() // 排序心愿项
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
      body: JSON.stringify({'content': content, 'creater': user})
    })
    if (!response.ok) {
      throw new Error('Failed to add wish')
    }
    const newWish = await response.json()
    wishes.value.data.push(newWish) // 先更新本地数据
    sortWishes() // 排序心愿项
  } catch (error) {
    console.error('Error adding wish:', error)
  }
  newWishContent.value = '' // 清空输入框
  fetchWishes()
}

// 删除心愿单
async function deleteWish(id) {
  // 二次询问
  if (!confirm('确定删除心愿吗？')) {
    return
  }
  try {
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
    const response = await fetch(`http://localhost:3051/api/update_wish/${id}`, {
      method: 'PUT',
      headers: {
        'Authorization': `${Cookies.get('token')}`,
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
      // 等待200ms
      await new Promise(resolve => setTimeout(resolve, 200))
      sortWishes() // 排序心愿项
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

// 排序心愿项
function sortWishes() {
  wishes.value.data.sort((a, b) => {
    if (!a.checked && !b.checked) {
      // 未勾选的心愿项按倒序排列
      return b.id - a.id; // 假设使用 id 进行倒序排列
    }
    if (a.checked && !b.checked) {
      // 已勾选的心愿项放在后面
      return 1;
    }
    if (!a.checked && b.checked) {
      // 未勾选的心愿项放在前面
      return -1;
    }
    // 已勾选的心愿项按顺序排列
    return a.id - b.id; // 假设使用 id 进行顺序排列
  });
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

.inputContent {
  text-align: center;
}

.wishContent {
  height: 70vh;
  width: 70%;
  border-radius: 20px;
}

input[type="checkbox"] {
  appearance: none;
  background-color: #ffffff00;
  margin: 0;
  font: inherit;
  color: #d9a67e;
  width: 2em;
  height: 2em;
  border: 0.3em solid currentColor;
  border-radius: 0.6em;
  display: grid;
  place-content: center;
  cursor: pointer;

}

input[type="checkbox"]::before {
  /* ...existing styles */
  background-color: #d9a67e;
  content: "";
  width: 1.1em;
  height: 1.1em;
  transform: scale(0);
  transition: 120ms transform ease-in-out;
  transform-origin: bottom left;
  clip-path: polygon(14% 44%, 0 65%, 50% 100%, 100% 16%, 80% 0%, 43% 62%);
}

input[type="checkbox"]:checked::before {
  transform: scale(1);
}
</style>
