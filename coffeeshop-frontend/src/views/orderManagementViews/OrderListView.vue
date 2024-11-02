<template>
    <div class="container">
        <h1>订单列表</h1>
        <n-data-table :columns="columns" :data="orders" :pagination="pagination" />
    </div>
</template>

<script setup>
import { ref, onMounted, h } from 'vue';
import { fetchOrders, updateOrderStatus } from '@/api/order'; // 修改为包含更新状态的函数
import { NButton } from "naive-ui";

const orders = ref([]);
const pagination = ref({ page: 1, pageSize: 10 }); // 分页配置

const columns = ref([
    { title: "订单 ID", key: "ID" },
    { title: "用户 ID", key: "UserID" },
    { title: "商品 ID", key: "ProductID" },
    { title: "数量", key: "Quantity" },
    { title: "价格", key: "Price" },
    { title: "订单状态", key: "Status" },
    {
        title: "操作",
        key: "actions",
        render(row) {
            const isDisabled = row.Status === "canceled" || row.Status === "completed";
            return [
                h(NButton, {
                    size: "small",
                    type: "success",
                    onClick: () => update(row.ID, "completed"),
                    disabled: isDisabled
                }, { default: () => "完成" }),
                h(NButton, {
                    size: "small",
                    type: "error",
                    onClick: () => update(row.ID, "canceled"),
                    disabled: isDisabled
                }, { default: () => "取消" })
            ];
        }
    }
]);

const fetchOrderList = async () => {
    try {
        const response = await fetchOrders();
        orders.value = response.data;
    } catch (error) {
        console.error("获取订单失败:", error);
    }
};

const update = async (id, status) => {
    try {
        await updateOrderStatus(id, { status }); // 调用更新状态的 API
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
