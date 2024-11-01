import axiosInstance from "../utils/axiosInstance";

// 获取用户列表请求
export const fetchUsers = () => {
  return axiosInstance.get("/users");
};
export const deleteUser = (id) => {
  return axiosInstance.delete(`/users/${id}`);
};
