import { createRouter, createWebHistory } from "vue-router";
import Loginview from "../views/loginview.vue";
import AdminView from "../views/adminView.vue";
import UserView from "../views/userView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "login",
      component: Loginview,
    },
    {
      path: "/admin",
      name: "admin",
      component: AdminView,
      meta: { requiresAuth: true, isAdmin: true }, // 需要认证且为管理员路由
    },
    {
      path: "/user",
      name: "user",
      component: UserView,
    },
  ],
});

router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem("isAuthenticated") === "true";
  const userRole = localStorage.getItem("userRole"); // 获取用户角色
  const requiresAuth = to.meta.requiresAuth;
  const isAdminRoute = to.meta.isAdmin; // 检查是否为管理员路由
  console.log(to);
  if (requiresAuth && !isAuthenticated) {
    next({ name: "login" }); // 未登录，重定向到登录页
  } else if (isAdminRoute && userRole !== "admin") {
    next({ name: "user" }); // 非管理员用户尝试访问管理员路由，重定向到主页
  } else {
    next(); // 允许访问
  }
});

export default router;
