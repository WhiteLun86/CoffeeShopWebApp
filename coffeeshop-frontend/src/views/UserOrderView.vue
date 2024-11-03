<template>
    <div class="container">
        <h1>我的订单</h1>
        <n-data-table :columns="columns" :data="orders" :pagination="pagination" />
    </div>
</template>

<script setup>
import { ref, onMounted, h } from 'vue';
import { fetchUserOrders } from '@/api/order'; // 引入获取用户订单的 API
import { NTag} from 'naive-ui';

const orders = ref([]);
const pagination = ref({ page: 1, pageSize: 10 });

// 订单列表的列配置
const columns = ref([
    { title: '订单 ID', key: 'id' },
    { title: '总价格', key: 'total_price' },
    { title: '订单状态', key: 'status' },
    {
        title: '商品信息',
        key: 'products',
        maxWidth: 500,
        render(row) {
            return row.products.map(product => 
                h(NTag, { key: product.ProductID }, `${product.product.Name} - 数量: ${product.quantity} - 价格: ¥${product.price}`)
            );
        }
    },
]);

const fetchOrderList = async () => {
    try {
        const response = await fetchUserOrders(parseInt(localStorage.getItem('userId'))); // 获取当前用户的订单
        orders.value = response.data;
    } catch (error) {
        console.error('获取订单失败:', error);
    }
};


// 页面加载时获取订单列表
onMounted(() => {
    fetchOrderList();
});
</script>

<style scoped>
.container {
    padding: 20px;
}
</style>

