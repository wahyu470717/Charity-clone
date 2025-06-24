import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

export interface Donation {
  id: number;
  amount: number;
  message?: string;
  donor_id: number;
  campaign_id: number;
  payment_status: 'pending' | 'completed' | 'failed';
  payment_method: string;
  created_at: string;
  updated_at: string;
  donor_name?: string;
  campaign_title?: string;
}

export const donationApi = {
  createDonation: async (donationData: {
    amount: number;
    message?: string;
    campaign_id: number;
    payment_method: string;
  }): Promise<Donation> => {
    const token = localStorage.getItem('token');
    const response = await axios.post(`${API_URL}/donations`, donationData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  getDonations: async (): Promise<Donation[]> => {
    const token = localStorage.getItem('token');
    const response = await axios.get(`${API_URL}/donations`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  },

  getCampaignDonations: async (campaignId: number): Promise<Donation[]> => {
    const response = await axios.get(`${API_URL}/campaigns/${campaignId}/donations`);
    return response.data;
  }
};