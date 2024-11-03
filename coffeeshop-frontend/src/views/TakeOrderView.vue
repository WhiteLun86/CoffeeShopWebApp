<template>
    <div class="container">
        <div class="product-list">
            <h2>商品列表</h2>
            <n-data-table :columns="productColumns" :data="products" />
        </div>

        <div class="cart">
            <h2>购物车</h2>
            <n-alert v-if="errorMessage" title="错误" type="error" closable @close="errorMessage = ''">
                {{ errorMessage }}
            </n-alert>
            <n-alert v-if="successMessage" title="成功" type="success" closable @close="successMessage = ''">
                {{ successMessage }}
            </n-alert>
            <n-data-table :columns="cartColumns" :data="cartItems" :pagination="pagination" />
            <div class="total-price">总价: ¥{{ totalPrice.toFixed(2) }}</div>
            <n-button type="primary" @click="placeOrder" :disabled="cartItems.length === 0">提交订单</n-button>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, h, computed } from 'vue';
import { fetchCoffeeProducts, } from '@/api/coffee'; // 引入获取商品的 API
import { NButton, createDiscreteApi } from "naive-ui";
import { order } from '@/api/order'

const { message } = createDiscreteApi(["message"]);
const products = ref([]);
const cartItems = ref([]);
const pagination = ref({ page: 1, pageSize: 10 });
const errorMessage = ref('');
const successMessage = ref('');
// 商品列表的列配置
const productColumns = ref([
    { title: 'ID', key: 'ID' },
    { title: '名称', key: 'Name' },
    { title: '价格', key: 'Price' },
    { title: '描述', key: 'Description' },
    { title: '库存', key: 'Stock' },
    {
        title: '操作',
        key: 'actions',
        render(row) {
            return h(NButton, {
                type: 'success',
                size: 'small',
                onClick: () => addToCart(row)
            }, { default: () => '加入购物车' });
        }
    }
]);

// 购物车的列配置
const cartColumns = ref([
    { title: '商品名称', key: 'name' },
    { title: '价格', key: 'price' },
    { title: '数量', key: 'quantity' },
    {
        title: '操作',
        key: 'actions',
        render(row) {
            return h(NButton, {
                type: 'error',
                size: 'small',
                onClick: () => removeFromCart(row)
            }, { default: () => '移除' });
        }
    }
]);

const fetchProductList = async () => {
    try {
        const response = await fetchCoffeeProducts();
        products.value = response.data;
    } catch (error) {
        console.error("获取商品失败:", error);
    }
};

// 添加商品到购物车
const addToCart = (product) => {
    const existingItem = cartItems.value.find(item => item.id === product.ID);
    if (existingItem) {
        existingItem.quantity++;
    } else {
        cartItems.value.push({ id: product.ID, name: product.Name, price: product.Price, quantity: 1 });
    }
};

// 从购物车移除商品
const removeFromCart = (item) => {
    cartItems.value = cartItems.value.filter(cartItem => cartItem.id !== item.id);
};

// 计算总价
const totalPrice = computed(() => {
    return cartItems.value.reduce((total, item) => total + item.price * item.quantity, 0);
});

// 提交订单

const placeOrder = async () => {
    if (cartItems.value.length === 0) {
        return; // 防止空购物车提交
    }

    // 准备订单数据
    const orderData = {
        user_id: parseInt(localStorage.getItem('userId')), // 假设用户 ID 从本地存储获取
        products: cartItems.value.map(item => ({
            product_id: item.id,
            quantity: item.quantity
        })),
        total_price: totalPrice.value
    };

    try {
        const response = await order(orderData);
        console.log('订单提交成功:', response);
        message.success('订单提交成功！');
    } catch (error) {
        console.error('订单提交失败:', error);
        // 检查后端返回的错误信息
        if (error.response && error.response.data && error.response.data.message) {
            message.error(error.response.data.message); // 显示后端返回的错误信息
        } else {
            message.error('订单提交失败，请稍后再试！');
        }
    }

    // 清空购物车
    cartItems.value = [];
};

// 页面加载时获取商品列表
onMounted(() => {
    fetchProductList();
});
</script>

<style scoped>
.container {
    padding: 20px;
}

.product-list,
.cart {
    margin-bottom: 20px;
}

.total-price {
    margin-top: 10px;
    font-weight: bold;
}
</style>
