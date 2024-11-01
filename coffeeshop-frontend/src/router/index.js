import { createRouter, createWebHistory } from "vue-router";
import Loginview from "../views/loginView.vue";
import AdminView from "../views/adminView.vue";
import UserView from "../views/userView.vue";
import AddUserView from "../views/userManagementViews/addUserView.vue";
import UserListView from "../views/userManagementViews/userListView.vue";
import AddProductView from "../views/productManagementViews/addproductView.vue";
import ProductListView from "../views/productManagementViews/productListView.vue";
import OrderListView from "../views/orderManagementViews/orderListView.vue";

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
      children: [
        {
          path: "add-user",
          name: "addUser",
          component: AddUserView,
          meta: { requiresAuth: true, isAdmin: true }, // 子路由继承meta
        },
        {
          path: "user-list",
          name: "userList",
          component: UserListView,
          meta: { requiresAuth: true, isAdmin: true },
        },
        {
          path: "product-list",
          name: "productList",
          component: ProductListView,
          meta: { requiresAuth: true, isAdmin: true },
        },
        {
          path: "add-product",
          name: "addProduct",
          component: AddProductView,
          meta: { requiresAuth: true, isAdmin: true }, 
        },
        {
          path: "order-list",
          name: "orderList",
          component: OrderListView,
          meta: { requiresAuth: true, isAdmin: true }, 
        },
      ],
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
