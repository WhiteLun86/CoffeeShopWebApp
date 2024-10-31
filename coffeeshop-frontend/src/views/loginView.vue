<template>
  <div class="login-container">
    <n-card title="☕ 登录 ☕" embedded class="login-card">
      <n-form @submit.prevent="handleLogin" label-align="left" label-width="80px">
        <n-form-item label="用户名">
          <n-input v-model:value="username" placeholder="请输入用户名" clearable />
        </n-form-item>
        <n-form-item label="密码">
          <n-input type="password" v-model:value="password" placeholder="请输入密码" clearable />
        </n-form-item>
        <n-form-item>
          <n-button type="primary" block attr-type="submit">登录</n-button>
        </n-form-item>
        <!-- 错误消息显示 -->
        <n-alert title="出错！" v-if="errorMessage" type="error" closable @close="errorMessage = ''">
          {{ errorMessage }}
        </n-alert>
        <n-alert title="成功!" v-if="successMessage" type="success" closable @close="successMessage = ''">
          {{ successMessage }}
        </n-alert>
      </n-form>
    </n-card>
  </div>
</template>

<script>
import { ref } from 'vue'
import { login } from '@/api/auth'
import { useRouter } from 'vue-router'
export default {
  setup() {
    const username = ref('')
    const password = ref('')
    const errorMessage = ref('')
    const successMessage = ref('')
    const router = useRouter() // 获取路由实例

    const handleLogin = async () => {
      try {
        const response = await login(username.value, password.value)
        console.log('登录成功:', response.data)
        // 显示成功提示
        errorMessage.value = ''
        successMessage.value = '登录成功！'

        // 延迟 1 秒后跳转到
        setTimeout(() => {
          if (response.data.role === 'admin') {
            router.push('/admin')
            localStorage.setItem('isAuthenticated', 'true');
            localStorage.setItem('userRole', 'admin');
          }else {
            router.push('/user')
            localStorage.setItem('isAuthenticated', 'true');
            localStorage.setItem('userRole', 'user');
          }
        }, 500)

      } catch (error) {
        if (error.response) {
          // 根据状态码显示相应的错误信息
          if (error.response.status === 404) {
            errorMessage.value = '用户不存在，请检查用户名。'
          } else if (error.response.status === 401) {
            errorMessage.value = '密码错误，请重试。'
          } else {
            errorMessage.value = '登录失败，请稍后重试。'
          }
        } else {
          errorMessage.value = '网络错误，请检查您的网络连接。'
        }
      }
    }

    return {
      username,
      password,
      handleLogin,
      errorMessage,
      successMessage
    }
  }
}
</script>

<style scoped>
.login-container {
  background-image: url('@/assets/images/login-bg.jpg');
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-size: cover;
  background-position: center;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 24px;
  background-color: rgba(255, 255, 255, 0.85);
  box-shadow: 0px 4px 14px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}
</style>