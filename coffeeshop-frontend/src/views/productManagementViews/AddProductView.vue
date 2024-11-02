<template>
    <div class="container">
        <n-card title="添加咖啡商品" embedded class="add-product-card">
            <n-form @submit.prevent="add" label-align="left" label-width="80px">
                <n-form-item label="商品名称">
                    <n-input v-model:value="product.name" placeholder="请输入商品名称" clearable />
                </n-form-item>
                <n-form-item label="价格">
                    <n-input-number v-model:value="product.price" placeholder="请输入商品价格" clearable :min="0" />
                </n-form-item>
                <n-form-item label="库存数量">
                    <n-input-number v-model:value="product.stock" placeholder="请输入商品库存数量" clearable :min="0" />
                </n-form-item>
                <n-form-item label="描述">
                    <n-input v-model:value="product.description" placeholder="请输入商品描述" clearable />
                </n-form-item>
                <n-form-item>
                    <n-button type="primary" block attr-type="submit">添加商品</n-button>
                </n-form-item>
                <n-alert v-if="errorMessage" title="错误" type="error" closable @close="errorMessage = ''">
                    {{ errorMessage }}
                </n-alert>
                <n-alert v-if="successMessage" title="成功" type="success" closable @close="successMessage = ''">
                    {{ successMessage }}
                </n-alert>
            </n-form>
        </n-card>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { addProduct } from '@/api/coffee'; // 确保路径正确

const product = ref({
    name: '',
    price: 0,
    stock: 0, // 新增库存数量
    description: ''
});
const errorMessage = ref('');
const successMessage = ref('');

const add = async () => {
    try {
        // 确保价格和库存数量是有效数字
        if (product.value.price <= 0) {
            errorMessage.value = '价格必须是正数！';
            return;
        }
        if (product.value.stock < 0) {
            errorMessage.value = '库存数量不能为负数！';
            return;
        }
        await addProduct(product.value);
        successMessage.value = '商品添加成功！';
        // 清空表单
        product.value = { name: '', price: 0, stock: 0, description: '' };
    } catch (error) {
        errorMessage.value = '添加商品失败，请稍后重试。';
        console.error("添加商品失败:", error);
    }
};
</script>

<style scoped>
.container {
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
}

.add-product-card {
    width: 100%;
    max-width: 400px;
    padding: 24px;
    background-color: rgba(255, 255, 255, 0.85);
    box-shadow: 0px 4px 14px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
}
</style>
