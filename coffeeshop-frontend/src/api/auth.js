import axiosInstance from '../utils/axiosInstance'

// 登录请求
export const login = (username, password) => {
  return axiosInstance.post('/login', { username, password })
}
