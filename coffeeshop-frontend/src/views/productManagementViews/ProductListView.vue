<template>
    <div>
        <h1>咖啡商品列表</h1>
        <n-data-table :single-line="false" :columns="columns" :data="data" />
    </div>
</template>

<script setup>
import { ref, h, onBeforeMount } from 'vue';
import { NButton, NForm, NFormItem, NInput, NInputNumber, } from "naive-ui";
import { fetchCoffeeProducts, deleteProduct, updateProduct } from '@/api/coffee'; 
import { createDiscreteApi } from "naive-ui";

const { message } = createDiscreteApi(["message"]);
const { dialog } = createDiscreteApi(["dialog"]);

let data = ref([]);
const columns = ref([
    {
        title: "ID",
        key: "ID"
    },
    {
        title: "名称",
        key: "Name"
    },
    {
        title: "价格",
        key: "Price"
    },
    {
        title: "库存数量",
        key: "Stock"
    },
    {
        title: "描述",
        key: "Description"
    },
    {
        title: "操作",
        key: "actions",
        render(row) {
            return [
                h(
                    NButton,
                    {
                        size: "small",
                        onClick: () => handleEdit(row)
                    },
                    { default: () => "编辑" }
                ),
                h(
                    NButton,
                    {
                        size: "small",
                        onClick: () => handleDelete(row)
                    },
                    { default: () => "删除" }
                )
            ];
        }
    }
]);

// 删除商品的函数
function handleDelete(row) {
    dialog.warning({
        title: "警告",
        content: "你确定要删除这个商品吗？",
        positiveText: "确定",
        negativeText: "不确定",
        onPositiveClick: async () => {
            try {
                await deleteProduct(row.ID);
                message.success('商品删除成功！');
                getProducts()
            } catch (error) {
                console.error("删除商品失败:", error);
                message.error('删除商品失败，请重试。');
            }
        },
        onNegativeClick: () => {
            message.warning("取消删除");
        }
    });
}

// 编辑商品信息
function handleEdit(row) {
    const name = ref(row.Name)
    const price = ref(row.Price)
    const stock = ref(row.Stock)
    const description = ref(row.Description)

    dialog.create({
        title: "编辑商品",
        content: () =>
            h(
                NForm,
                { labelWidth: "80px" },
                {
                    default: () => [
                        h(
                            NFormItem,
                            { label: "名称" },
                            () =>
                                h(NInput, {
                                    value: name.value,
                                    onUpdateValue: (value) => (name.value = value),
                                    clearable: true,
                                })
                        ),
                        h(
                            NFormItem,
                            { label: "价格" },
                            () =>
                                h(NInputNumber, {
                                    value: price.value,
                                    onUpdateValue: (value) => (price.value = value),
                                    clearable: true,
                                    min: 0
                                })
                        ),
                        h(
                            NFormItem,
                            { label: "库存数量" },
                            () =>
                                h(NInputNumber, {
                                    value: stock.value,
                                    onUpdateValue: (value) => (stock.value = value),
                                    clearable: true,
                                    min: 0
                                })
                        ),
                        h(
                            NFormItem,
                            { label: "描述" },
                            () =>
                                h(NInput, {
                                    value: description.value,
                                    onUpdateValue: (value) => (description.value = value),
                                    clearable: true
                                })
                        )
                    ]
                }
            ),
        positiveText: "保存",
        negativeText: "取消",
        onPositiveClick: async () => {
            try {
                // 调用后端接口更新商品信息
                await updateProduct({
                    ID: row.ID,
                    Name: name.value,
                    Stock: stock.value,
                    Description: description.value
                });
                message.success('商品信息更新成功！');
                getProducts(); // 刷新商品列表
            } catch (error) {
                console.error('商品信息更新失败:', error);
                message.error('商品信息更新失败，请重试。');
            }
        }
    });
}

const getProducts = async () => {
    try {
        const response = await fetchCoffeeProducts();
        data.value = response.data; // 将获取的商品数据赋值给 data
    } catch (error) {
        console.error("获取商品数据失败:", error);
    }
};

// 在组件挂载时获取商品列表
onBeforeMount(() => {
    getProducts();
});
</script>

<style scoped>
/* 样式 */
</style>
