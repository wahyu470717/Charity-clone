import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { 
  Container, 
  Box, 
  Typography, 
  Button, 
  Card, 
  CardContent,
  LinearProgress,
  Chip,
  Avatar,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Stack
} from '@mui/material';
import { 
  // Person,
   CalendarToday, 
   AttachMoney } from '@mui/icons-material';
import { useQuery } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import { donationApi } from '../../api/donation';
import Loading from '../../components/common/Loading';
import { useAuth } from '../../contexts/AuthContext';

const CampaignDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { user } = useAuth();
  
  const { data: campaign, isLoading: campaignLoading } = useQuery({
    queryKey: ['campaign', id],
    queryFn: () => campaignApi.getCampaign(Number(id)),
    enabled: !!id
  });

  const { data: donations, isLoading: donationsLoading } = useQuery({
    queryKey: ['campaign-donations', id],
    queryFn: () => donationApi.getCampaignDonations(Number(id)),
    enabled: !!id
  });

  if (campaignLoading) return <Loading />;
  
  if (!campaign) {
    return (
      <Container>
        <Typography variant="h6">Campaign not found</Typography>
      </Container>
    );
  }

  const progress = (campaign.current_amount / campaign.target_amount) * 100;
  const daysLeft = Math.ceil((new Date(campaign.end_date).getTime() - new Date().getTime()) / (1000 * 60 * 60 * 24));

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      {/* Menggunakan CSS Grid Layout untuk responsive design */}
      <Box
        sx={{
          display: 'grid',
          gridTemplateColumns: {
            xs: '1fr', // Mobile: 1 column
            md: '2fr 1fr' // Desktop: 2:1 ratio
          },
          gap: 4
        }}
      >
        {/* Main Content Area */}
        <Box>
          <Box mb={3}>
            <img 
              src={campaign.image_url || '/placeholder-campaign.jpg'} 
              alt={campaign.title}
              style={{ width: '100%', height: '300px', objectFit: 'cover', borderRadius: '8px' }}
            />
          </Box>
          
          <Typography variant="h4" gutterBottom>
            {campaign.title}
          </Typography>
          
          <Stack direction="row" spacing={2} mb={3}>
            <Chip 
              label={campaign.status} 
              color={campaign.status === 'active' ? 'success' : 'default'}
            />
            <Chip 
              icon={<CalendarToday />}
              label={`${daysLeft > 0 ? `${daysLeft} days left` : 'Ended'}`}
            />
          </Stack>
          
          <Typography variant="body1" paragraph>
            {campaign.description}
          </Typography>
        </Box>
        
        {/* Sidebar */}
        <Stack spacing={3}>
          <Card>
            <CardContent>
              <Box mb={2}>
                <Typography variant="h6" color="primary">
                  ${campaign.current_amount.toLocaleString()}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  raised of ${campaign.target_amount.toLocaleString()} goal
                </Typography>
              </Box>
              
              <LinearProgress 
                variant="determinate" 
                value={Math.min(progress, 100)} 
                sx={{ height: 8, borderRadius: 4, mb: 2 }}
              />
              
              <Typography variant="body2" color="text.secondary" mb={3}>
                {progress.toFixed(1)}% of goal reached
              </Typography>
              
              {user && campaign.status === 'active' && (
                <Button 
                  variant="contained" 
                  fullWidth 
                  size="large"
                  onClick={() => navigate(`/campaigns/${campaign.id}/donate`)}
                >
                  Donate Now
                </Button>
              )}
              
              {!user && (
                <Button 
                  variant="contained" 
                  fullWidth 
                  size="large"
                  onClick={() => navigate('/login')}
                >
                  Login to Donate
                </Button>
              )}
            </CardContent>
          </Card>
          
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Recent Donations
              </Typography>
              {donationsLoading ? (
                <Loading />
              ) : (
                <List dense>
                  {donations?.slice(0, 5).map((donation) => (
                    <ListItem key={donation.id} sx={{ px: 0 }}>
                      <ListItemAvatar>
                        <Avatar>
                          <AttachMoney />
                        </Avatar>
                      </ListItemAvatar>
                      <ListItemText
                        primary={`$${donation.amount}`}
                        secondary={donation.donor_name || 'Anonymous'}
                      />
                    </ListItem>
                  ))}
                </List>
              )}
            </CardContent>
          </Card>
        </Stack>
      </Box>
    </Container>
  );
};

export default CampaignDetailPage;