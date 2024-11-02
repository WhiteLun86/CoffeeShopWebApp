import axiosInstance from "../utils/axiosInstance";

export const fetchUsers = () => {
  return axiosInstance.get("/users");
};

export const deleteUser = (id) => {
  return axiosInstance.delete(`/users/${id}`);
};

export const updateUser = (userData) => {
  return axiosInstance.put('/user/update', userData)
}
export const createUser = (userData) => {
  return axiosInstance.post('/users', userData)
}