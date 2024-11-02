// src/api/order.js
import axiosInstance from "../utils/axiosInstance";

// 获取订单列表
export const fetchOrders = () => {
  return axiosInstance.get("/orders");
};

// 删除订单
export const deleteOrder = (id) => {
  return axiosInstance.delete(`/orders/${id}`);
};

// 删除订单
export const updateOrderStatus = (id, status) => {
  return axiosInstance.put(`/orders/${id}/status`, status );
};
