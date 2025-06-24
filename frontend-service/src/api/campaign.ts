import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

export interface Campaign {
  id: number;
  title: string;
  description: string;
  short_description: string;
  target_amount: number;
  current_amount: number;
  start_date: string;
  end_date: string;
  image_url?: string;
  status: string;
  recipient_id: number;
  created_by_id: number;
  created_at: string;
  updated_at: string;
  recipient_name?: string;
  created_by_name?: string;
}

export const campaignApi = {
  getCampaigns: async (): Promise<Campaign[]> => {
    const response = await axios.get(`${API_URL}/campaigns`);
    return response.data;
  },

  getCampaign: async (id: number): Promise<Campaign> => {
    const response = await axios.get(`${API_URL}/campaigns/${id}`);
    return response.data;
  },

  createCampaign: async (campaignData: {
    title: string;
    description: string;
    short_description: string;
    target_amount: number;
    start_date: string;
    end_date: string;
    image_url?: string;
    recipient_id: number;
  }): Promise<Campaign> => {
    const token = localStorage.getItem('token');
    const response = await axios.post(`${API_URL}/campaigns`, campaignData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  updateCampaign: async (id: number, campaignData: Partial<Campaign>): Promise<Campaign> => {
    const token = localStorage.getItem('token');
    const response = await axios.put(`${API_URL}/campaigns/${id}`, campaignData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  deleteCampaign: async (id: number): Promise<void> => {
    const token = localStorage.getItem('token');
    await axios.delete(`${API_URL}/campaigns/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  }
};