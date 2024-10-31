// src/utils/axiosInstance.js
import axios from 'axios'

// 创建 axios 实例
const instance = axios.create({
  baseURL: '/api', // 后端API的基础URL，根据实际情况调整
  timeout: 5000, // 请求超时设置
})

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 可添加 token 或其他请求头信息
    // config.headers.Authorization = `Bearer ${token}`
    return config
  },
  error => Promise.reject(error)
)

// 响应拦截器
instance.interceptors.response.use(
  response => response,
  error => {
    console.error('请求出错：', error)
    return Promise.reject(error)
  }
)

export default instance
