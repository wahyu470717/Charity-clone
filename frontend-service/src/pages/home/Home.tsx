import React from 'react';
import { 
  Container, 
  Box, 
  Typography, 
  Button, 
  Card,
  CardContent,
  CardMedia,
  Chip,
  Stack,
  useTheme, 
  useMediaQuery
} from '@mui/material';
import Grid from '@mui/material/Grid';

import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import CampaignCard from '../../components/campaign/CampaignCard';

const articles = [
  {
    id: 1,
    category: 'Nonprofit',
    readTime: '5 min read',
    title: 'The Power of Community Giving',
    description: 'Learn how community efforts can change lives for the better.',
    image: 'https://placehold.co/200x120',
  },
  {
    id: 2,
    category: 'Charity',
    readTime: '5 min read',
    title: 'Innovative Fundraising Strategies for Nonprofits',
    description: 'Discover new ways to engage donors and maximize contributions.',
    image: 'https://placehold.co/200x120',
  },
  {
    id: 3,
    category: 'CC',
    readTime: '5 min read',
    title: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
    description: 'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    image: 'https://placehold.co/200x120',
  },
];

// const quotes = [
//   {
//     name: 'Emily Johnson',
//     position: 'Survivor, Community Advocate',
//     avatar: 'https://placehold.co/80x80',
//     rating: 5,
//     text: '“Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse varius enim in eros elementum tristique.”',
//   },
//   {
//     name: 'Name Surname',
//     position: 'Position, Company name',
//     avatar: 'https://placehold.co/80x80',
//     rating: 5,
//     text: '“Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut commodo diam libero vitae erat.”',
//   },
//   {
//     name: 'Name Surname',
//     position: 'Position, Company name',
//     avatar: 'https://placehold.co/80x80',
//     rating: 5,
//     text: '“Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet. Nunc ut sem vitae risus.”',
//   },
//   {
//     name: 'Name Surname',
//     position: 'Position, Company name',
//     avatar: 'https://placehold.co/80x80',
//     rating: 5,
//     text: '“Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.”',
//   },
// ];

const HomePage: React.FC = () => {
  const { data: campaigns } = useQuery({
    queryKey: ['campaigns'],
    queryFn: campaignApi.getCampaigns
  });

  const featuredCampaigns = campaigns?.slice(0, 3) || [];

  const theme = useTheme();
  const isMdUp = useMediaQuery(theme.breakpoints.up('md'));

  return (
    <>
      <Box sx={{ bgcolor: '#0a0a23', py: { xs: 8, md: 12 }, px: 2 }}>
        <Container maxWidth="xl">
          <Box
            sx={{
              display: 'flex',
              flexDirection: { xs: 'column', md: 'row' },
              alignItems: 'center',
              justifyContent: 'space-between',
              gap: { xs: 4, md: 8 },
            }}
          >
            {/* Left: Text Content */}
            <Box flex={1} sx={{ color: 'white' }}>
              <Box
                sx={{
                  width: 48,
                  height: 48,
                  borderRadius: '50%',
                  bgcolor: 'rgba(255,255,255,0.1)',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  mb: 2,
                }}
              >
                {/* Replace this with a proper icon */}
                <img src="https://placehold.co/40x40" alt="icon" width={24} height={24} />
              </Box>
              <Typography variant={isMdUp ? 'h4' : 'h5'} fontWeight="bold" gutterBottom>
                Discover the Heart of Giving with Heals Mealt
              </Typography>
              <Typography variant="body1" sx={{ color: '#ccc', mb: 3 }}>
                Heals Mealt uniquely blends donation initiatives with engaging blog content, 
                fostering a vibrant community of supporters. Our platform not only raises 
                awareness but also inspires action, making a real difference in the world.
              </Typography>
              <Stack direction="row" spacing={2} flexWrap="wrap">
                <Button variant="contained" color="secondary" component={Link} to="/campaigns" sx={{ borderRadius: 5 }}>
                  Donate
                </Button>
                <Button variant="text" color="inherit" component={Link} to="/about" sx={{ borderRadius: 5 }}>
                  Learn More &gt;
                </Button>
              </Stack>
            </Box>

            {/* Right: Image */}
            <Box flex={1} display="flex" justifyContent="center">
              <Card
                elevation={0}
                sx={{
                  borderRadius: 3,
                  overflow: 'hidden',
                  maxWidth: '100%',
                  width: { xs: '100%', sm: '80%', md: '100%' },
                }}
              >
                <img
                  src="https://placehold.co/800x450"
                  alt="Giving"
                  style={{
                    width: '100%',
                    height: 'auto',
                    display: 'block',
                    objectFit: 'cover',
                  }}
                />
              </Card>
            </Box>
          </Box>
        </Container>
      </Box>

      <Box sx={{ flexGrow: 1, bgcolor: '#000', py: 10, px: 2, color: 'white' }}>
        <Container maxWidth="lg">
          {/* Header */}
          <Box textAlign="center" mb={6}>
            <Typography variant="h4" fontWeight="bold" gutterBottom>
              Explore Our Latest Insights
            </Typography>
            <Typography variant="body1" sx={{ color: '#aaa' }}>
              Discover impactful stories and initiatives in giving.
            </Typography>
          </Box>

          <Grid container spacing={{ xs: 2, md: 3 }} columns={{ xs: 4, sm: 8, md: 12 }}>
            {articles.map((article) => (
              <Grid size={{ xs: 2, sm: 4, md: 4 }} key={article.id}>
                <Card 
                  sx={{
                    bgcolor: '#111',
                    color: 'white',
                    borderRadius: 3,
                    overflow: 'hidden',
                    boxShadow: 'none',
                    height: '100%',
                  }}
                >
                  <CardMedia
                    component="img"
                    image={article.image}
                    alt={article.title}
                    sx={{
                      height: 200,
                      width: '100%',
                      objectFit: 'cover',
                    }}
                  />
                  <CardContent>
                    <Stack direction="row" spacing={1} alignItems="center" mb={1}>
                      <Chip label={article.category} size="small" sx={{ fontSize: '0.75rem' }} />
                      <Typography variant="caption" sx={{ color: '#aaa' }}>
                        {article.readTime}
                      </Typography>
                    </Stack>
                    <Typography variant="h6" fontWeight="bold" gutterBottom>
                      {article.title}
                    </Typography>
                    <Typography variant="body2" sx={{ color: '#ccc', mb: 2 }}>
                      {article.description}
                    </Typography>
                    <Typography
                      variant="body2"
                      sx={{
                        fontWeight: 500,
                        color: 'white',
                        cursor: 'pointer',
                        '&:hover': { textDecoration: 'underline' },
                      }}
                    >
                      Read more &gt;
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>

          {/* Featured Campaigns */}
          <Container maxWidth="xl" sx={{ py: 3 }}>
    
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
            
            {/* View All Button */}
            <Box textAlign="center">
              <Button
                variant="outlined"
                sx={{
                  borderRadius: '999px',
                  color: 'white',
                  borderColor: '#444',
                  textTransform: 'none',
                  fontWeight: 'medium',
                  '&:hover': {
                    borderColor: 'white',
                  },
                }}
                component={Link} 
                to="/campaigns"
              >
                View all
              </Button>
            </Box>
          </Container>

        </Container>
      </Box>

    </>
  );
};

export default HomePage;