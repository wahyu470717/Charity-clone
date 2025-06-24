import React, { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { 
  Container, 
  Box, 
  Typography, 
  TextField, 
  Button, 
  Card, 
  CardContent,
  FormControl,
  FormLabel,
  RadioGroup,
  FormControlLabel,
  Radio,
  Alert,
  Stack
} from '@mui/material';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { campaignApi } from '../../api/campaign';
import { donationApi } from '../../api/donation';
import Loading from '../../components/common/Loading';

const DonatePage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  
  const [amount, setAmount] = useState('');
  const [customAmount, setCustomAmount] = useState('');
  const [message, setMessage] = useState('');
  const [paymentMethod, setPaymentMethod] = useState('credit_card');
  
  const { data: campaign, isLoading } = useQuery({
    queryKey: ['campaign', id],
    queryFn: () => campaignApi.getCampaign(Number(id)),
    enabled: !!id
  });

  const donationMutation = useMutation({
    mutationFn: donationApi.createDonation,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['campaign', id] });
      navigate(`/campaigns/${id}`, { 
        state: { message: 'Thank you for your donation!' } 
      });
    },
  });

  const handleAmountSelect = (value: string) => {
    setAmount(value);
    setCustomAmount('');
  };

  const handleCustomAmountChange = (value: string) => {
    setCustomAmount(value);
    setAmount('');
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const donationAmount = amount ? Number(amount) : Number(customAmount);
    
    if (donationAmount <= 0) {
      return;
    }
    
    donationMutation.mutate({
      amount: donationAmount,
      message: message.trim() || undefined,
      campaign_id: Number(id),
      payment_method: paymentMethod
    });
  };

  if (isLoading) return <Loading />;
  
  if (!campaign) {
    return (
      <Container>
        <Typography variant="h6">Campaign not found</Typography>
      </Container>
    );
  }

  const finalAmount = amount ? Number(amount) : Number(customAmount) || 0;

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      <Typography variant="h4" gutterBottom>
        Donate to {campaign.title}
      </Typography>
      
      {/* Main layout using Stack */}
      <Stack 
        direction={{ xs: 'column', md: 'row' }}
        spacing={4}
        sx={{ alignItems: 'flex-start' }}
      >
        {/* Main Donation Form */}
        <Card sx={{ flex: { md: 2 } }}>
          <CardContent>
            <Box component="form" onSubmit={handleSubmit}>
              <Typography variant="h6" gutterBottom>
                Choose Amount
              </Typography>
              
              {/* Amount selection buttons using CSS Grid */}
              <Box
                sx={{
                  display: 'grid',
                  gridTemplateColumns: {
                    xs: 'repeat(2, 1fr)', // Mobile: 2 columns
                    sm: 'repeat(4, 1fr)' // Desktop: 4 columns
                  },
                  gap: 2,
                  mb: 3
                }}
              >
                {['10', '25', '50', '100'].map((value) => (
                  <Button
                    key={value}
                    variant={amount === value ? 'contained' : 'outlined'}
                    onClick={() => handleAmountSelect(value)}
                  >
                    ${value}
                  </Button>
                ))}
              </Box>
              
              <TextField
                fullWidth
                label="Custom Amount"
                type="number"
                value={customAmount}
                onChange={(e) => handleCustomAmountChange(e.target.value)}
                sx={{ mb: 3 }}
                InputProps={{
                  startAdornment: <Typography>$</Typography>
                }}
              />
              
              <TextField
                fullWidth
                label="Message (Optional)"
                multiline
                rows={3}
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                sx={{ mb: 3 }}
              />
              
              <FormControl component="fieldset" sx={{ mb: 3 }}>
                <FormLabel component="legend">Payment Method</FormLabel>
                <RadioGroup
                  value={paymentMethod}
                  onChange={(e) => setPaymentMethod(e.target.value)}
                >
                  <FormControlLabel 
                    value="credit_card" 
                    control={<Radio />} 
                    label="Credit Card" 
                  />
                  <FormControlLabel 
                    value="bank_transfer" 
                    control={<Radio />} 
                    label="Bank Transfer" 
                  />
                  <FormControlLabel 
                    value="paypal" 
                    control={<Radio />} 
                    label="PayPal" 
                  />
                </RadioGroup>
              </FormControl>
              
              {donationMutation.error && (
                <Alert severity="error" sx={{ mb: 2 }}>
                  {donationMutation.error instanceof Error 
                    ? donationMutation.error.message 
                    : 'Donation failed'}
                </Alert>
              )}
              
              <Button
                type="submit"
                variant="contained"
                size="large"
                fullWidth
                disabled={finalAmount <= 0 || donationMutation.isPending}
              >
                {donationMutation.isPending ? 'Processing...' : `Donate $${finalAmount}`}
              </Button>
            </Box>
          </CardContent>
        </Card>
        
        {/* Campaign Summary Sidebar */}
        <Card sx={{ flex: { md: 1 } }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Campaign Summary
            </Typography>
            <Typography variant="body2" paragraph>
              {campaign.short_description}
            </Typography>
            <Stack spacing={1}>
              <Typography variant="body2" color="text.secondary">
                Goal: ${campaign.target_amount.toLocaleString()}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Raised: ${campaign.current_amount.toLocaleString()}
              </Typography>
            </Stack>
          </CardContent>
        </Card>
      </Stack>
    </Container>
  );
};

export default DonatePage;