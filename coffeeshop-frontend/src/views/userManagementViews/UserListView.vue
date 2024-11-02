<template>
    <div> <h1>咖啡商品列表</h1>
        <n-data-table :single-line="false" :columns="columns" :data="data" /></div>
</template>

<script setup>
import { ref, h, onBeforeMount } from 'vue';
import { NSelect, NButton, NForm, NFormItem, NInput } from "naive-ui";
import { fetchUsers, deleteUser, updateUser } from '@/api/user';
import { createDiscreteApi } from "naive-ui"
const { message } = createDiscreteApi(["message"])
const { dialog } = createDiscreteApi(["dialog"])
function handleDelete(row) {
    dialog.warning({
        title: "警告",
        content: "你确定？",
        positiveText: "确定",
        negativeText: "不确定",
        onPositiveClick: async () => {
            try {
                await deleteUser(row.id);
                message.success('用户删除成功！');
                getUsers(); // 刷新用户列表
            } catch (error) {
                console.error("删除用户失败:", error);
                message.error('删除用户失败，请重试。');
            }

            message.success("确定");
        },
        onNegativeClick: () => {
            message.warning("不确定");
        }
    });
}
let data = ref([]);
const columns = ref([
    {
        title: "Id",
        key: "id"
    },
    {
        title: "用户名",
        key: "username"
    },
    {
        title: "密码",
        key: "password"
    },
    {
        title: "用户类型",
        key: "role",
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
                        onClick: () => handleDelete(row) // 改为删除用户的函数
                    },
                    { default: () => "删除" }
                )
            ];
        }
    }
]);


// 编辑用户信息
function handleEdit(row) {
    const editedUsername = ref(row.username)
    const editedPassword = ref('')
    const editedRole = ref(row.role)

    dialog.create({
        title: "编辑用户",
        content: () =>
            h(
                NForm,
                { labelWidth: "80px" },
                {
                    default: () => [
                        h(
                            NFormItem,
                            { label: "用户名" },
                            () =>
                                h(NInput, {
                                    value: editedUsername.value,
                                    onUpdateValue: (value) => (editedUsername.value = value),
                                    clearable: true
                                })
                        ),
                        h(
                            NFormItem,
                            { label: "密码" },
                            () =>
                                h(NInput, {
                                    type: "password",
                                    value: editedPassword.value,
                                    onUpdateValue: (value) => (editedPassword.value = value),
                                    clearable: true
                                })
                        ),
                        h(
                            NFormItem,
                            { label: "用户类型" },
                            () =>
                                h(NSelect, {
                                    options: [
                                        { label: "User", value: "user" },
                                        { label: "Admin", value: "admin" }
                                    ],
                                    value: editedRole.value,
                                    onUpdateValue: (value) => (editedRole.value = value),
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
                // 调用后端接口更新用户信息
                await updateUser({
                    id: row.id,
                    username: editedUsername.value,
                    password: editedPassword.value,
                    role: editedRole.value
                })
                message.success('用户信息更新成功！')
                setTimeout(getUsers(), 500)
            } catch (error) {
                console.error('用户信息更新失败:', error)
                message.error('用户信息更新失败，请重试。')
            }
        }
    })
}

const getUsers = async () => {
    try {
        const response = await fetchUsers();
        data.value = response.data; // 将获取的用户数据赋值给 data
        console.log(data);
    } catch (error) {
        console.error("获取用户数据失败:", error);
    }
};

// 在组件挂载时获取用户列表
onBeforeMount(() => {
    getUsers();
});
</script>

<style scoped>
/* 样式 */
</style>
