<template>
    <div class="register-container">
        <n-card title="☕ 注册 ☕" embedded class="register-card">
            <n-form @submit.prevent="handleRegister" label-align="left" label-width="80px">
                <n-form-item label="用户名">
                    <n-input v-model:value="username" placeholder="请输入用户名" clearable />
                </n-form-item>
                <n-form-item label="密码">
                    <n-input type="password" v-model:value="password" placeholder="请输入密码" clearable />
                </n-form-item>
                <n-form-item label="确认密码">
                    <n-input type="password" v-model:value="confirmPassword" placeholder="请确认密码" clearable />
                </n-form-item>
                <n-form-item class="button-container">
                    <n-button type="primary" round attr-type="submit">注册</n-button>
                    <n-button type="primary" round @click="goToLogin">返回登录</n-button>
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
import { createUser } from '@/api/user' // 假设有一个注册的 API
import { useRouter } from 'vue-router'

export default {
    setup() {
        const username = ref('')
        const password = ref('')
        const confirmPassword = ref('')
        const errorMessage = ref('')
        const successMessage = ref('')
        const router = useRouter() // 获取路由实例

        const handleRegister = async () => {
            if (password.value !== confirmPassword.value) {
                errorMessage.value = '密码不匹配，请确认后重试。'
                return
            }

            try {
                const response = await await createUser({
                    username: username.value,
                    password: password.value,
                    role: 'user'
                })
                console.log('注册成功:', response.data)
                successMessage.value = '注册成功！'
                errorMessage.value = ''

                // 延迟 1 秒后跳转到登录页面
                setTimeout(() => {
                    router.push('/')
                }, 1000)

            } catch (error) {
                if (error.response) {
                    errorMessage.value = '注册失败，请稍后重试。'
                } else {
                    errorMessage.value = '网络错误，请检查您的网络连接。'
                }
            }
        }

        // 跳转到登录页面
        const goToLogin = () => {
            router.push('/'); // 确保 '/login' 是你的登录页面的路由
        }

        return {
            username,
            password,
            confirmPassword,
            handleRegister,
            goToLogin,
            errorMessage,
            successMessage
        }
    }
}
</script>

<style scoped>
.button-container {
    display: flex;
    flex-direction: column;
}

.register-container {
    background-image: url('@/assets/images/login-bg.jpg');
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    background-size: cover;
    background-position: center;
}

.register-card {
    width: 100%;
    max-width: 400px;
    padding: 24px;
    background-color: rgba(255, 255, 255, 0.85);
    box-shadow: 0px 4px 14px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
}
</style>