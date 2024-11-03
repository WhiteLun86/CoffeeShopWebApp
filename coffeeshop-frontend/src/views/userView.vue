<template>
    <div class="main-container">
        <n-flex vertical size="large">
            <n-layout has-sider>
                <n-layout-sider bordered collapse-mode="width" :collapsed-width="64" :width="240" :collapsed="collapsed"
                    show-trigger @collapse="collapsed = true" @expand="collapsed = false">
                    <n-menu accordion :collapsed="collapsed" :collapsed-width="64" :collapsed-icon-size="22"
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
import { ref, h, onMounted } from 'vue';
import { RouterView, RouterLink, useRouter } from 'vue-router';
const route = useRouter();
const collapsed = ref(false);
onMounted(() => {
    route.replace('/admin');
})
const menuOptions = ref([
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "order",
                }
            },
            { default: () => "下单商品" }
        ),
        key: "order",
    },
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "userorder",
                }
            },
            { default: () => "订单管理" }
        ),
        key: "userorder",
    },
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