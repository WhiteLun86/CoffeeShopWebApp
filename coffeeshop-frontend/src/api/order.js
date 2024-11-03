import axiosInstance from "../utils/axiosInstance";

export const fetchOrders = () => {
  return axiosInstance.get("/orders");
};
export const deleteOrder = (id) => {
  return axiosInstance.delete(`/orders/${id}`);
};
export const updateOrderStatus = (id, status) => {
  return axiosInstance.put(`/orders/${id}/status`, status );
};
export const order = (orderData) => {
  return axiosInstance.post('/orders', orderData);
};
export const fetchUserOrders = (userId) => {
  return axiosInstance.get(`/user/${userId}/orders`);
};
