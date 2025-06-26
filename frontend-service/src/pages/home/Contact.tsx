import React from 'react';
import { Container, Typography,
  //  Box, 
   Card, CardContent, Stack } from '@mui/material';
import { Email, Phone, LocationOn } from '@mui/icons-material';

const ContactPage: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      <Typography variant="h3" component="h1" gutterBottom textAlign="center">
        Contact Us
      </Typography>
      
      <Typography variant="body1" textAlign="center" paragraph>
        Get in touch with our team for support, partnerships, or any questions.
      </Typography>

      {/* Alternative using Stack - simpler approach */}
      <Stack 
        direction={{ xs: 'column', sm: 'row' }}
        spacing={4}
        sx={{
          flexWrap: { sm: 'wrap', md: 'nowrap' },
          justifyContent: 'center',
          alignItems: 'stretch'
        }}
      >
        <Card sx={{ textAlign: 'center', flex: { sm: '1 1 45%', md: '1' } }}>
          <CardContent>
            <Email sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Email
            </Typography>
            <Typography variant="body2">
              support@sharethemeal.com
            </Typography>
          </CardContent>
        </Card>
        
        <Card sx={{ textAlign: 'center', flex: { sm: '1 1 45%', md: '1' } }}>
          <CardContent>
            <Phone sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Phone
            </Typography>
            <Typography variant="body2">
              +1 (555) 123-4567
            </Typography>
          </CardContent>
        </Card>
        
        <Card sx={{ textAlign: 'center', flex: { sm: '1 1 100%', md: '1' } }}>
          <CardContent>
            <LocationOn sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Address
            </Typography>
            <Typography variant="body2">
              123 Charity Street<br />
              City, State 12345
            </Typography>
          </CardContent>
        </Card>
      </Stack>
    </Container>
  );
};

export default ContactPage;