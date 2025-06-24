import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { API_BASE_URL, API_TIMEOUT } from '../constants';
import { ApiResponse, PaginationParams } from '../types';
import { authService } from './authService';

class ApiService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: API_BASE_URL,
      timeout: API_TIMEOUT,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor
    this.api.interceptors.request.use(
      (config) => {
        const token = authService.getAuthToken();
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // Response interceptor
    this.api.interceptors.response.use(
      (response) => response.data,
      (error) => {
        if (error.response?.status === 401) {
          authService.clearAuthToken();
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  // Generic CRUD methods
  async get<T>(endpoint: string, params?: any): Promise<ApiResponse<T>> {
    const config: AxiosRequestConfig = {};
    if (params) {
      config.params = params;
    }
    return this.api.get(endpoint, config);
  }

  async post<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    return this.api.post(endpoint, data);
  }

  async put<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    return this.api.put(endpoint, data);
  }

  async patch<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    return this.api.patch(endpoint, data);
  }

  async delete<T>(endpoint: string): Promise<ApiResponse<T>> {
    return this.api.delete(endpoint);
  }

  // File upload
  async upload<T>(endpoint: string, file: File, data?: any): Promise<ApiResponse<T>> {
    const formData = new FormData();
    formData.append('file', file);
    
    if (data) {
      Object.keys(data).forEach(key => {
        formData.append(key, data[key]);
      });
    }

    return this.api.post(endpoint, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
  }

  // Multiple file upload
  async uploadMultiple<T>(endpoint: string, files: File[], data?: any): Promise<ApiResponse<T>> {
    const formData = new FormData();
    
    files.forEach((file, index) => {
      formData.append('files', file);
    });
    
    if (data) {
      Object.keys(data).forEach(key => {
        formData.append(key, data[key]);
      });
    }

    return this.api.post(endpoint, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
  }

  // Paginated requests
  async getPaginated<T>(endpoint: string, params: PaginationParams): Promise<ApiResponse<T[]>> {
    return this.get(endpoint, params);
  }

  // Export data
  async export(endpoint: string, format: 'csv' | 'excel' | 'pdf', params?: any): Promise<Blob> {
    const response = await this.api.get(`${endpoint}/export`, {
      params: { ...params, format },
      responseType: 'blob',
    });
  return response as unknown as Blob;
  }

  // Bulk operations
  async bulkCreate<T>(endpoint: string, items: any[]): Promise<ApiResponse<T[]>> {
    return this.post(`${endpoint}/bulk`, { items });
  }

  async bulkUpdate<T>(endpoint: string, items: any[]): Promise<ApiResponse<T[]>> {
    return this.put(`${endpoint}/bulk`, { items });
  }

  async bulkDelete<T>(endpoint: string, ids: string[]): Promise<ApiResponse<T>> {
    return this.delete(`${endpoint}/bulk?ids=${ids.join(',')}`);
  }

  // Search
  async search<T>(endpoint: string, query: string, filters?: any): Promise<ApiResponse<T[]>> {
    return this.get(`${endpoint}/search`, { q: query, ...filters });
  }

  // Statistics
  async getStats<T>(endpoint: string, params?: any): Promise<ApiResponse<T>> {
    return this.get(`${endpoint}/stats`, params);
  }
}

export const apiService = new ApiService();