<template>
    <div class="create-user-container">
      <n-card title="创建用户" embedded class="create-user-card">
        <n-form @submit.prevent="handleCreateUser" label-align="left" label-width="80px">
          <n-form-item label="用户名" path="username" >
            <n-input v-model:value="username" placeholder="请输入用户名"  clearable minlength="3"/>
          </n-form-item>
          <n-form-item label="密码" path="password">
            <n-input type="password" v-model:value="password" placeholder="请输入密码"  clearable minlength="6"/>
          </n-form-item>
          <n-form-item label="用户类型" path="role">
            <n-select v-model:value="role" :options="roleOptions" placeholder="请选择用户类型" />
          </n-form-item>
          <n-form-item>
            <n-button type="primary" block attr-type="submit">创建用户</n-button>
          </n-form-item>
        </n-form>
      </n-card>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue'
  import { createUser } from '@/api/user'
  import { createDiscreteApi } from 'naive-ui'
  
  const { message } = createDiscreteApi(['message'])
  
  export default {
    setup() {
      const username = ref('')
      const password = ref('')
      const role = ref('')
      const roleOptions = [
        { label: 'User', value: 'user' },
        { label: 'Admin', value: 'admin' }
      ]
  
      const handleCreateUser = async () => {
        if (!username.value || !password.value || !role.value) {
          message.error('请填写所有必填项')
          return
        }
  
        try {
          await createUser({
            username: username.value,
            password: password.value,
            role: role.value
          })
          message.success('用户创建成功！')
          // 清空表单
          username.value = ''
          password.value = ''
          role.value = ''
        } catch (error) {
          console.error('用户创建失败:', error)
          message.error('用户创建失败，请重试')
        }
      }
  
      return {
        username,
        password,
        role,
        roleOptions,
        handleCreateUser
      }
    }
  }
  </script>
  
  <style scoped>
  .create-user-container {
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }
  .create-user-card {
    width: 100%;
    max-width: 400px;
    padding: 24px;
    background-color: rgba(255, 255, 255, 0.85);
    box-shadow: 0px 4px 14px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
  }
  </style>
  