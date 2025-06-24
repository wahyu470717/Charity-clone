import React from 'react';
import { Box, Grid, Typography } from '@mui/material';
import { useQuery } from 'react-query';
import { getDashboardStats } from '@/services/dashboardService';
import StatCard from '@/components/dashboard/StatCard';
import ActivityFeed from '@/components/dashboard/ActivityFeed';
import RevenueChart from '@/components/dashboard/RevenueChart';

const DashboardPage: React.FC = () => {
  const { data: stats, isLoading, error } = useQuery('dashboardStats', getDashboardStats);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading dashboard data</div>;

  return (
    <Box>
      <Typography variant="h4" gutterBottom>Dashboard</Typography>
      
      <Grid container spacing={3} mb={4}>
        <Grid item xs={12} sm={6} md={3}>
          <StatCard 
            title="Total Campaigns" 
            value={stats?.totalCampaigns || 0} 
            icon="campaign" 
            color="#4caf50"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <StatCard 
            title="Active Campaigns" 
            value={stats?.activeCampaigns || 0} 
            icon="active_campaign" 
            color="#2196f3"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <StatCard 
            title="Total Donations" 
            value={stats?.totalDonations || 0} 
            icon="volunteer_activism" 
            color="#ff9800"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <StatCard 
            title="Total Amount" 
            value={stats?.totalAmount || 0} 
            icon="payments" 
            color="#9c27b0"
            isCurrency
          />
        </Grid>
      </Grid>

      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <RevenueChart data={stats?.revenueData || []} />
        </Grid>
        <Grid item xs={12} md={4}>
          <ActivityFeed activities={stats?.recentActivities || []} />
        </Grid>
      </Grid>
    </Box>
  );
};

export default DashboardPage;