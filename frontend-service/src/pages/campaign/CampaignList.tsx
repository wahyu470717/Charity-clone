import React, { useState } from 'react';
import { 
  Container, 
  Typography, 
  Box,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button,
  Stack
} from '@mui/material';
import type { SelectChangeEvent } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';
import CampaignCard from '../../components/campaign/CampaignCard';
import { useQuery } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import Loading from '../../components/common/Loading';
import { useAuth } from '../../contexts/AuthContext';
import { Link } from 'react-router-dom';

const CampaignListPage: React.FC = () => {
  const { user } = useAuth();
  const [searchTerm, setSearchTerm] = useState('');
  const [statusFilter, setStatusFilter] = useState('all');
  
  const { data: campaigns, isLoading, error } = useQuery({
    queryKey: ['campaigns'],
    queryFn: campaignApi.getCampaigns
  });

  const handleStatusChange = (e: SelectChangeEvent<string>) => {
    setStatusFilter(e.target.value);
  };

  const filteredCampaigns = campaigns?.filter(campaign => {
    const matchesSearch = campaign.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         campaign.short_description.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesStatus = statusFilter === 'all' || campaign.status === statusFilter;
    return matchesSearch && matchesStatus;
  });

  if (isLoading) return <Loading />;
  
  if (error) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Typography variant="h6" color="error">
          Error loading campaigns. Please try again later.
        </Typography>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h4" component="h1">
          Active Campaigns
        </Typography>
        {user && (user.role === 'recipient' || user.role === 'superadmin') && (
          <Button
            variant="contained"
            startIcon={<AddIcon />}
            component={Link}
            to="/campaigns/create"
          >
            Create Campaign
          </Button>
        )}
      </Box>

      <Box mb={3}>
        <Stack 
          direction={{ xs: 'column', md: 'row' }} 
          spacing={2} 
          alignItems="stretch"
        >
          <Box flex={2}>
            <TextField
              fullWidth
              label="Search campaigns"
              variant="outlined"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </Box>
          <Box flex={1} minWidth={200}>
            <FormControl fullWidth>
              <InputLabel>Status</InputLabel>
              <Select
                value={statusFilter}
                label="Status"
                onChange={handleStatusChange}
              >
                <MenuItem value="all">All</MenuItem>
                <MenuItem value="active">Active</MenuItem>
                <MenuItem value="completed">Completed</MenuItem>
                <MenuItem value="cancelled">Cancelled</MenuItem>
              </Select>
            </FormControl>
          </Box>
        </Stack>
      </Box>

      <Box
        sx={{
          display: 'grid',
          gridTemplateColumns: {
            xs: '1fr',
            sm: 'repeat(2, 1fr)',
            md: 'repeat(3, 1fr)',
          },
          gap: 3,
        }}
      >
        {filteredCampaigns?.map((campaign) => (
          <Box key={campaign.id}>
            <CampaignCard campaign={campaign} />
          </Box>
        ))}
      </Box>

      {filteredCampaigns?.length === 0 && (
        <Box textAlign="center" mt={4}>
          <Typography variant="h6" color="text.secondary">
            No campaigns found matching your criteria.
          </Typography>
        </Box>
      )}
    </Container>
  );
};

export default CampaignListPage;