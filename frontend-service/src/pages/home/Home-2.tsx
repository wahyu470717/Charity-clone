import React from 'react';
import { 
  Container, 
  Box, 
  Typography, 
  Button, 
  Card, 
  CardContent,
  // CardMedia,
  Stack
} from '@mui/material';
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import CampaignCard from '../../components/campaign/CampaignCard';

const HomePage: React.FC = () => {
  const { data: campaigns } = useQuery({
    queryKey: ['campaigns'],
    queryFn: campaignApi.getCampaigns
  });

  const featuredCampaigns = campaigns?.slice(0, 3) || [];

  return (
    <Box>
      {/* Hero Section */}
      <Box
        sx={{
          background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
          color: 'white',
          py: 8,
          textAlign: 'center',
          width: '100vw',
        }}
      >
        <Container maxWidth="lg">
          <Typography variant="h2" component="h1" gutterBottom>
            Share The Meal
          </Typography>
          <Typography variant="h5" component="p" gutterBottom>
            Help those in need by donating to food campaigns around the world
          </Typography>
          <Stack 
            direction={{ xs: 'column', sm: 'row' }} 
            spacing={2} 
            justifyContent="center"
            sx={{ mt: 4 }}
          >
            <Button 
              variant="contained" 
              size="large" 
              component={Link} 
              to="/campaigns"
              sx={{ bgcolor: 'white', color: 'primary.main' }}
            >
              Browse Campaigns
            </Button>
            <Button 
              variant="outlined" 
              size="large" 
              component={Link} 
              to="/about"
              sx={{ borderColor: 'white', color: 'white' }}
            >
              Learn More
            </Button>
          </Stack>
        </Container>
      </Box>

      {/* Featured Campaigns */}
      <Container maxWidth="lg" sx={{ py: 8 }}>
        <Typography variant="h4" component="h2" textAlign="center" gutterBottom>
          Featured Campaigns
        </Typography>
        <Typography variant="body1" textAlign="center" color="text.secondary" mb={4}>
          Support these urgent food campaigns
        </Typography>
        
        {/* Alternative using Stack for featured campaigns */}
        <Stack 
          direction={{ xs: 'column', sm: 'row' }}
          spacing={4}
          sx={{ 
            flexWrap: { sm: 'wrap', md: 'nowrap' },
            justifyContent: 'center',
            alignItems: 'stretch',
            mb: 4
          }}
        >
          {featuredCampaigns.map((campaign) => (
            <Box key={campaign.id} sx={{ flex: { sm: '1 1 45%', md: '1' } }}>
              <CampaignCard campaign={campaign} />
            </Box>
          ))}
        </Stack>
        
        <Box textAlign="center">
          <Button variant="contained" component={Link} to="/campaigns">
            View All Campaigns
          </Button>
        </Box>
      </Container>

      {/* How It Works */}
      <Box sx={{ bgcolor: 'grey.50', py: 8 }}>
        <Container maxWidth="lg">
          <Typography variant="h4" component="h2" textAlign="center" gutterBottom>
            How It Works
          </Typography>
          
          {/* Alternative using Stack for how it works cards */}
          <Stack 
            direction={{ xs: 'column', sm: 'row' }}
            spacing={4}
            sx={{ 
              flexWrap: { sm: 'wrap', md: 'nowrap' },
              justifyContent: 'center',
              alignItems: 'stretch',
              mt: 4
            }}
          >
            <Card sx={{ textAlign: 'center', p: 2, flex: { sm: '1 1 45%', md: '1' } }}>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  1. Browse Campaigns
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Find food campaigns that need your support
                </Typography>
              </CardContent>
            </Card>
            
            <Card sx={{ textAlign: 'center', p: 2, flex: { sm: '1 1 45%', md: '1' } }}>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  2. Make a Donation
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Choose an amount and donate securely
                </Typography>
              </CardContent>
            </Card>
            
            <Card sx={{ textAlign: 'center', p: 2, flex: { sm: '1 1 100%', md: '1' } }}>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  3. Make an Impact
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Help provide meals to those in need
                </Typography>
              </CardContent>
            </Card>
          </Stack>
        </Container>
      </Box>
    </Box>
  );
};

export default HomePage;