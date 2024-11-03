<template>
    <div class="container">
        <h1>订单列表</h1>
        <n-data-table :columns="columns" :data="orders" :pagination="pagination" />
    </div>
</template>

<script setup>
import { ref, onMounted, h } from 'vue';
import { fetchOrders, updateOrderStatus } from '@/api/order'; 
import { NButton, NTag } from "naive-ui";

const orders = ref([]);
const pagination = ref({ page: 1, pageSize: 10 }); // 分页配置

const columns = ref([
    { title: "订单 ID", key: "id" },
    { title: "用户 ID", key: "user_id" },
    { 
        title: "商品信息", 
        key: "products", 
        render(row) {
            return row.products.map(product => 
                h(NTag, { key: product.ProductID }, `${product.product.Name} - 数量: ${product.quantity} - 价格: ¥${product.price}`)
            );
        },
        maxWidth: 500
    },
    { title: "总价格", key: "total_price" },
    { title: "订单状态", key: "status" },
    {
        title: "操作",
        key: "actions",
        render(row) {
            const isDisabled = row.status === "canceled" || row.status === "completed";
            return [
                h(NButton, {
                    size: "small",
                    type: "success",
                    onClick: () => update(row.id, "completed"),
                    disabled: isDisabled
                }, { default: () => "完成" }),
                h(NButton, {
                    size: "small",
                    type: "error",
                    onClick: () => update(row.id, "canceled"),
                    disabled: isDisabled
                }, { default: () => "取消" })
            ];
        }
    }
]);

const fetchOrderList = async () => {
    try {
        const response = await fetchOrders();
        orders.value = response.data; // 直接赋值
    } catch (error) {
        console.error("获取订单失败:", error);
    }
};

const update = async (id, status) => {
    try {
        await updateOrderStatus(id, { status });
        await fetchOrderList(); // 重新获取订单列表
    } catch (error) {
        console.error("更新订单状态失败:", error);
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
