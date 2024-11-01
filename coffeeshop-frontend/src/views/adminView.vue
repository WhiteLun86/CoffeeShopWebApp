<template>
    <div class="main-container">
        <n-flex vertical size="large">
            <n-layout has-sider>
                <n-layout-sider bordered collapse-mode="width" :collapsed-width="64" :width="240" :collapsed="collapsed"
                    show-trigger @collapse="collapsed = true" @expand="collapsed = false">
                    <n-menu :collapsed="collapsed" :collapsed-width="64" :collapsed-icon-size="22"
                        :options="menuOptions" />
                </n-layout-sider>
                <n-layout-content content-style="padding: 24px;">
                    <RouterView></RouterView>
                </n-layout-content>
            </n-layout>
        </n-flex>
    </div>

</template>
<script setup>
import { ref, h } from 'vue';
import { RouterView, RouterLink } from 'vue-router';
const collapsed = ref(false);
const menuOptions = ref([
    {
        label: "用户管理",
        key: "user-manage",
        children: [
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: "userList",
                        }
                    },
                    { default: () => "用户列表" }
                ),
                key: "user-list"
            },
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: "addUser",
                        }
                    },
                    { default: () => "创建用户" }
                ),
                key: "user-create"
            },

        ]
    },
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "orderList",
                }
            },
            { default: () => "订单管理" }
        ),
        key: "order-manage",
    },
    {
        label: "库存管理",
        key: "product-manage",
        children: [
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: "productList",
                        }
                    },
                    { default: () => "商品列表" }
                ),
                key: "product-list"
            },
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: "addProduct",
                        }
                    },
                    { default: () => "添加商品" }
                ),
                key: "product-create"
            },

        ]
    }
])
</script>
<style scoped>
.main-container {
    width: 100vw;
    height: 100vh;
}

.n-flex {
    height: 100%;
}


.n-layout-sider {
    background: rgba(128, 128, 128, 0.3);
}

.n-layout-content {
    background-image: url('@/assets/images/main-bg.png');
    background-size: cover;
    /* 使背景图片覆盖整个区域 */
    background-position: center;
    /* 居中背景图片 */
}
</style>