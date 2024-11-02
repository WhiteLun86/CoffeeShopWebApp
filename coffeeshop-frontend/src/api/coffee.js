import axiosInstance from '../utils/axiosInstance';

// 获取咖啡商品列表的请求
export const fetchCoffeeProducts = () => {
  return axiosInstance.get('/products'); // 假设后端路由为 /coffee-products
};
export const deleteProduct = (id) => {
    return axiosInstance.delete(`/products/${id}`);
}
export const addProduct = (product) => {
    return axiosInstance.post('/products', product);
}
export const updateProduct = async (product) => {
  return axiosInstance.put(`/coffee/${product.ID}`, product);
};