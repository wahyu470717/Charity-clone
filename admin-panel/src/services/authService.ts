import axios, { AxiosInstance } from 'axios';
import { API_BASE_URL, API_TIMEOUT } from '../constants';
import { AuthResponse, AuthUser, LoginCredentials } from '../types';

class AuthService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: `${API_BASE_URL}/auth`,
      timeout: API_TIMEOUT,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor to add auth token
    this.api.interceptors.request.use(
      (config) => {
        const token = this.getAuthToken();
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => {
        return Promise.reject(error);
      }
    );

    // Response interceptor for error handling
    this.api.interceptors.response.use(
      (response) => response.data,
      (error) => {
        if (error.response?.status === 401) {
          // Token expired, redirect to login
          this.clearAuthToken();
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  private authToken: string = '';

  setAuthToken(token: string): void {
    this.authToken = token;
  }

  getAuthToken(): string {
    return this.authToken;
  }

  clearAuthToken(): void {
    this.authToken = '';
  }

  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    const response = await this.api.post('/login', credentials);
    return response.data;
  }

  async logout(): Promise<void> {
    await this.api.post('/logout');
  }

  async refreshToken(refreshToken: string): Promise<AuthResponse> {
    const response = await this.api.post('/refresh', { refreshToken });
    return response.data;
  }

  async verifyToken(): Promise<{ user: AuthUser }> {
    const response = await this.api.get('/verify');
    return response.data;
  }

  async forgotPassword(email: string): Promise<{ message: string }> {
    const response = await this.api.post('/forgot-password', { email });
    return response.data;
  }

  async resetPassword(token: string, password: string): Promise<{ message: string }> {
    const response = await this.api.post('/reset-password', { token, password });
    return response.data;
  }

  async changePassword(currentPassword: string, newPassword: string): Promise<{ message: string }> {
    const response = await this.api.post('/change-password', {
      currentPassword,
      newPassword,
    });
    return response.data;
  }

  async updateProfile(userData: Partial<AuthUser>): Promise<{ user: AuthUser }> {
    const response = await this.api.put('/profile', userData);
    return response.data;
  }

  async uploadAvatar(file: File): Promise<{ avatarUrl: string }> {
    const formData = new FormData();
    formData.append('avatar', file);

    const response = await this.api.post('/upload-avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  }
}

export const authService = new AuthService();