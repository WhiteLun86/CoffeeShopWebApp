<template>
  <div
    class="login-container"
  >
    <n-card title="☕ 登录 ☕" embedded class="login-card">
      <n-form @submit.prevent="handleLogin" label-align="left" label-width="80px">
        <n-form-item label="用户名">
          <n-input v-model="username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item label="密码">
          <n-input type="password" v-model="password" placeholder="请输入密码" />
        </n-form-item>
        <n-form-item>
          <n-button type="primary" native-type="submit" block>登录</n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  setup() {
    const username = ref('')
    const password = ref('')

    const handleLogin = () => {
      console.log('用户名:', username.value)
      console.log('密码:', password.value)
    }

    return {
      username,
      password,
      handleLogin,
    }
  },
}
</script>

<style scoped>
/* 背景图片和容器样式 */
.login-container {
  background-image: url('@/assets/images/login-bg.jpg');
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-size: cover;
  background-position: center;
}

/* 登录框样式 */
.login-card {
  width: 100%;
  max-width: 400px;
  padding: 24px;
  background-color: rgba(255, 255, 255, 0.85); /* 半透明背景 */
  box-shadow: 0px 4px 14px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}
</style>
