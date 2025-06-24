import React from 'react';
import { Container, Typography, Box, Card, CardContent, Stack } from '@mui/material';

const AboutPage: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      <Typography variant="h3" component="h1" gutterBottom textAlign="center">
        About Share The Meal
      </Typography>
      
      <Box sx={{ mb: 4 }}>
        <Typography variant="h5" gutterBottom>
          Our Mission
        </Typography>
        <Typography variant="body1" paragraph>
          Share The Meal is dedicated to fighting hunger by connecting generous donors 
          with communities in need. Our platform makes it easy to create and support 
          food campaigns that directly impact people's lives.
        </Typography>
      </Box>

      {/* Alternative using Stack - simpler approach */}
      <Stack 
        direction={{ xs: 'column', md: 'row' }} 
        spacing={4}
      >
        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              For Donors
            </Typography>
            <Typography variant="body2">
              Browse active campaigns, make secure donations, and track the impact 
              of your contributions in real-time.
            </Typography>
          </CardContent>
        </Card>
        
        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              For Recipients
            </Typography>
            <Typography variant="body2">
              Create campaigns to raise funds for food initiatives in your community 
              and reach donors who want to help.
            </Typography>
          </CardContent>
        </Card>
      </Stack>
    </Container>
  );
};

export default AboutPage;