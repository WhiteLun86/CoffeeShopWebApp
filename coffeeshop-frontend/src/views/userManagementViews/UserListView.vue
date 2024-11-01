<template>
    <n-dialog-provider>
        <n-data-table :single-line="false" :columns="columns" :data="data" />
    </n-dialog-provider>

</template>

<script setup>
import { ref, h, onBeforeMount } from 'vue';
import { NButton, NForm, NFormItem, NInput } from "naive-ui";
import { fetchUsers, deleteUser } from '@/api/user';
import { createDiscreteApi } from "naive-ui"
const { message } = createDiscreteApi(["message"])
const { dialog } = createDiscreteApi(["dialog"])
const editedUser = ref({}); // 用于存储当前编辑的用户数据
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


// 编辑用户函数
function handleEdit(row) {
    editedUser.value = { ...row }; // 将用户信息复制到 editedUser 以便进行编辑
    dialog.create({
        title: "编辑用户",
        content: () => [
            h(NForm, {
                model: editedUser.value,
                onSubmit: async () => {
                    try {
                        // await updateUser(editedUser.value);
                        message.success("用户信息更新成功！");
                        getUsers(); // 刷新用户列表
                    } catch (error) {
                        console.error("更新用户信息失败:", error);
                        message.error("更新失败，请重试。");
                    }
                }
            }, {
                default: () => [
                    h(NFormItem, { label: "用户名" }, () =>
                        h(NInput, {
                            value: editedUser.value.username,
                            onUpdateValue: (value) => editedUser.value.username = value
                        })
                    ),
                    h(NFormItem, { label: "密码" }, () =>
                        h(NInput, {
                            type: "password",
                            value: editedUser.value.password,
                            onUpdateValue: (value) => editedUser.value.password = value
                        })
                    ),
                    h(NFormItem, { label: "用户类型" }, () =>
                        h(NInput, {
                            value: editedUser.value.role,
                            onUpdateValue: (value) => editedUser.value.role = value
                        })
                    ),
                    h(
                        NButton,
                        { type: "primary", attrType: "submit", style: "margin-top: 10px" },
                        { default: () => "提交修改" }
                    )
                ]
            })
        ]
    });
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
