import axios from 'axios';

const API_URL ='http://localhost:8080/api/v1';

export interface User {
  id: number;
  name: string;
  email: string;
  role: 'donor' | 'recipient' | 'superadmin';
  profile_picture?: string;
  phone_number?: string;
  address?: string;
  created_at: string;
  updated_at: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export const authApi = {
  login: async (email: string, password: string): Promise<LoginResponse> => {
    const response = await axios.post(`${API_URL}/auth/login`, { email, password });
    return response.data;
  },

  register: async (userData: {
    name: string;
    email: string;
    password: string;
    role?: string;
  }): Promise<User> => {
    const response = await axios.post(`${API_URL}/auth/register`, userData);
    return response.data;
  },

  getCurrentUser: async (): Promise<User> => {
    const token = localStorage.getItem('token');
    const response = await axios.get(`${API_URL}/users/me`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  updateProfile: async (userData: {
    name?: string;
    profile_picture?: string;
    phone_number?: string;
    address?: string;
  }): Promise<User> => {
    const token = localStorage.getItem('token');
    const response = await axios.put(`${API_URL}/users/me`, userData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  updatePassword: async (currentPassword: string, newPassword: string): Promise<void> => {
    const token = localStorage.getItem('token');
    await axios.put(
      `${API_URL}/users/me/password`,
      { currentPassword, newPassword },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
  }
};
