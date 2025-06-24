import React from 'react';
import { 
  Card, 
  CardMedia, 
  CardContent, 
  Typography, 
  CardActions, 
  Button, 
  Box,
  LinearProgress,
  Chip,
  Stack
} from '@mui/material';
import { Link } from 'react-router-dom';
import type { Campaign } from '../../api/campaign';

interface CampaignCardProps {
  campaign: Campaign;
}

const CampaignCard: React.FC<CampaignCardProps> = ({ campaign }) => {
  const progress = (campaign.current_amount / campaign.target_amount) * 100;
  const daysLeft = Math.ceil((new Date(campaign.end_date).getTime() - new Date().getTime()) / (1000 * 60 * 60 * 24));

  return (
    <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <CardMedia
        component="img"
        height="200"
        image={campaign.image_url || '/placeholder-campaign.jpg'}
        alt={campaign.title}
      />
      <CardContent sx={{ flexGrow: 1 }}>
        <Stack 
          direction="row" 
          justifyContent="space-between" 
          alignItems="center" 
          sx={{ mb: 1 }}
        >
          <Chip 
            label={campaign.status} 
            color={campaign.status === 'active' ? 'success' : 'default'}
            size="small"
          />
          <Typography variant="body2" color="text.secondary">
            {daysLeft > 0 ? `${daysLeft} days left` : 'Ended'}
          </Typography>
        </Stack>
        
        <Typography gutterBottom variant="h6" component="div" noWrap>
          {campaign.title}
        </Typography>
        
        <Typography variant="body2" color="text.secondary" sx={{ mb: 2, height: 40, overflow: 'hidden' }}>
          {campaign.short_description}
        </Typography>
        
        <Box sx={{ mb: 2 }}>
          <Stack direction="row" justifyContent="space-between" sx={{ mb: 1 }}>
            <Typography variant="body2" color="text.secondary">
              Raised: ${campaign.current_amount.toLocaleString()}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {progress.toFixed(1)}%
            </Typography>
          </Stack>
          <LinearProgress 
            variant="determinate" 
            value={Math.min(progress, 100)} 
            sx={{ height: 8, borderRadius: 4 }}
          />
          <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5 }}>
            Goal: ${campaign.target_amount.toLocaleString()}
          </Typography>
        </Box>
      </CardContent>
      
      <CardActions sx={{ p: 2, pt: 0 }}>
        <Stack direction="row" spacing={1} sx={{ width: '100%' }}>
          <Button size="small" component={Link} to={`/campaigns/${campaign.id}`} fullWidth>
            Learn More
          </Button>
          <Button 
            size="small" 
            variant="contained" 
            component={Link} 
            to={`/campaigns/${campaign.id}/donate`}
            fullWidth
            disabled={campaign.status !== 'active'}
          >
            Donate
          </Button>
        </Stack>
      </CardActions>
    </Card>
  );
};

export default CampaignCard;