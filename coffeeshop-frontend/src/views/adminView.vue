<template>
    <div class="main-container">
        <n-flex vertical size="large">
            <n-layout has-sider>
                <n-layout-sider content-style="padding: 24px;">
                    <n-anchor>
                        <n-anchor-link title="演示" href="#演示">
                            <n-anchor-link title="基础用法" href="#basic.vue" />
                            <n-anchor-link title="忽略间隔" href="#ignore-gap.vue" />
                            <n-anchor-link title="固定" href="#affix.vue" />
                            <n-anchor-link title="滚动到" href="#scrollto.vue" />
                        </n-anchor-link>
                        <n-anchor-link title="API" href="#API" />
                    </n-anchor>
                </n-layout-sider>
                <n-layout>
                    <n-layout-header>颐和园路</n-layout-header>
                    <n-layout-content content-style="padding: 24px;">
                        平山道
                    </n-layout-content>
                </n-layout>
            </n-layout>
        </n-flex>
    </div>

</template>
<style scoped>
.main-container {
    width: 100vw;
    height: 100vh;
}

.n-flex {
    height: 100%;
}

.n-layout-header,
.n-layout-footer {
    background: rgba(128, 128, 128, 0.2);
    padding: 24px;
}

.n-layout-sider {
    background: rgba(128, 128, 128, 0.3);
}

.n-layout-content {
    background: rgba(128, 128, 128, 0.4);
}
</style>