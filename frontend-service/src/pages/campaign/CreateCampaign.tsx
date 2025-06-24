import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { 
  Container, 
  Box, 
  Typography, 
  TextField, 
  Button, 
  Card, 
  CardContent,
  Alert,
  Stack,
  InputAdornment
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers';
import { useMutation } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import { AxiosError } from 'axios';

interface CampaignFormData {
  title: string;
  short_description: string;
  description: string;
  target_amount: string;
  start_date: Date | null;
  end_date: Date | null;
  image_url: string;
}

const CreateCampaignPage: React.FC = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState<CampaignFormData>({
    title: '',
    short_description: '',
    description: '',
    target_amount: '',
    start_date: null,
    end_date: null,
    image_url: ''
  });
  const [error, setError] = useState('');

  const campaignMutation = useMutation({
    mutationFn: campaignApi.createCampaign,
    onSuccess: (data) => {
      navigate(`/campaigns/${data.id}`);
    },
    onError: (err: AxiosError<{ message?: string }>) => {
      setError(err.response?.data?.message || 'Failed to create campaign');
    }
  });

  const handleChange = (field: keyof CampaignFormData) => 
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setFormData(prev => ({ ...prev, [field]: e.target.value }));
    };

  const handleDateChange = (field: 'start_date' | 'end_date') => 
    (date: Date | null) => {
      setFormData(prev => ({ ...prev, [field]: date }));
    };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!formData.start_date || !formData.end_date) {
      setError('Please select start and end dates');
      return;
    }
    
    if (new Date(formData.end_date) <= new Date(formData.start_date)) {
      setError('End date must be after start date');
      return;
    }
    
    campaignMutation.mutate({
      title: formData.title,
      short_description: formData.short_description,
      description: formData.description,
      target_amount: Number(formData.target_amount),
      start_date: formData.start_date.toISOString().split('T')[0],
      end_date: formData.end_date.toISOString().split('T')[0],
      image_url: formData.image_url || undefined,
      recipient_id: 1 // Ini harus diganti dengan recipient_id yang sesuai
    });
  };

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      <Typography variant="h4" gutterBottom>
        Create New Campaign
      </Typography>
      
      <Card>
        <CardContent>
          <Box component="form" onSubmit={handleSubmit}>
            {error && <Alert severity="error" sx={{ mb: 3 }}>{error}</Alert>}
            
            <Stack spacing={3}>
              <TextField
                label="Campaign Title"
                fullWidth
                required
                value={formData.title}
                onChange={handleChange('title')}
              />
              
              <TextField
                label="Short Description"
                fullWidth
                required
                value={formData.short_description}
                onChange={handleChange('short_description')}
                helperText="A brief summary of your campaign (max 150 characters)"
                inputProps={{ maxLength: 150 }}
              />
              
              <TextField
                label="Full Description"
                fullWidth
                required
                multiline
                rows={6}
                value={formData.description}
                onChange={handleChange('description')}
              />
              
              <TextField
                label="Target Amount ($)"
                fullWidth
                required
                type="number"
                value={formData.target_amount}
                onChange={handleChange('target_amount')}
                InputProps={{
                  startAdornment: <InputAdornment position="start">$</InputAdornment>,
                }}
              />
              
              <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
                <DatePicker
                  label="Start Date"
                  value={formData.start_date}
                  onChange={handleDateChange('start_date')}
                  sx={{ flex: 1 }}
                />
                
                <DatePicker
                  label="End Date"
                  value={formData.end_date}
                  onChange={handleDateChange('end_date')}
                  sx={{ flex: 1 }}
                />
              </Stack>
              
              <TextField
                label="Image URL (Optional)"
                fullWidth
                value={formData.image_url}
                onChange={handleChange('image_url')}
                helperText="URL of an image for your campaign"
              />
              
              <Box sx={{ display: 'flex', justifyContent: 'flex-end', pt: 2 }}>
                <Button 
                  type="submit" 
                  variant="contained" 
                  size="large"
                  disabled={campaignMutation.isPending}
                >
                  {campaignMutation.isPending ? 'Creating...' : 'Create Campaign'}
                </Button>
              </Box>
            </Stack>
          </Box>
        </CardContent>
      </Card>
    </Container>
  );
};

export default CreateCampaignPage;