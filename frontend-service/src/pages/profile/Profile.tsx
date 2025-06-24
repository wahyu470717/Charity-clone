import React, { useState } from 'react';
import { 
  Container, 
  Box, 
  Typography, 
  Avatar, 
  TextField, 
  Button, 
  Card, 
  CardContent,
  Divider,
  Alert,
  Stack
} from '@mui/material';
import { useAuth } from '../../contexts/AuthContext';
import { authApi } from '../../api/auth';

const ProfilePage: React.FC = () => {
  const { user, updateUser } = useAuth();
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({
    name: user?.name || '',
    email: user?.email || '',
    phone_number: user?.phone_number || '',
    address: user?.address || '',
  });
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (field: string) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData(prev => ({ ...prev, [field]: e.target.value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (!user) return;
      
      const updatedUser = await authApi.updateProfile({
        name: formData.name,
        phone_number: formData.phone_number,
        address: formData.address
      });
      
      updateUser(updatedUser);
      setSuccess('Profile updated successfully');
      setEditMode(false);
      setTimeout(() => setSuccess(''), 3000);
    } catch (err) {
      setError('Failed to update profile. Please try again.');
      return err
    }
  };

  if (!user) {
    return (
      <Container>
        <Typography variant="h6">Please login to view your profile</Typography>
      </Container>
    );
  }

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      <Typography variant="h4" gutterBottom>
        My Profile
      </Typography>
      
      <Card>
        <CardContent>
          <Stack direction="row" spacing={3} alignItems="center" sx={{ mb: 3 }}>
            <Avatar
              alt={user.name}
              src={user.profile_picture}
              sx={{ width: 80, height: 80 }}
            />
            <Box>
              <Typography variant="h5">{user.name}</Typography>
              <Typography variant="body2" color="text.secondary">
                {user.role}
              </Typography>
            </Box>
          </Stack>

          {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
          {success && <Alert severity="success" sx={{ mb: 2 }}>{success}</Alert>}

          {editMode ? (
            <Box component="form" onSubmit={handleSubmit}>
              <Stack spacing={3}>
                <TextField
                  label="Full Name"
                  fullWidth
                  value={formData.name}
                  onChange={handleChange('name')}
                  required
                />
                
                <TextField
                  label="Email"
                  fullWidth
                  value={formData.email}
                  disabled
                />
                
                <TextField
                  label="Phone Number"
                  fullWidth
                  value={formData.phone_number}
                  onChange={handleChange('phone_number')}
                />
                
                <TextField
                  label="Address"
                  fullWidth
                  multiline
                  rows={3}
                  value={formData.address}
                  onChange={handleChange('address')}
                />
                
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button 
                    variant="outlined" 
                    onClick={() => setEditMode(false)}
                  >
                    Cancel
                  </Button>
                  <Button 
                    type="submit" 
                    variant="contained"
                  >
                    Save Changes
                  </Button>
                </Stack>
              </Stack>
            </Box>
          ) : (
            <Stack spacing={2}>
              <Box>
                <Typography variant="subtitle2" color="text.secondary">
                  Email
                </Typography>
                <Typography>{user.email}</Typography>
              </Box>
              
              <Box>
                <Typography variant="subtitle2" color="text.secondary">
                  Phone Number
                </Typography>
                <Typography>{user.phone_number || '-'}</Typography>
              </Box>
              
              <Box>
                <Typography variant="subtitle2" color="text.secondary">
                  Address
                </Typography>
                <Typography>{user.address || '-'}</Typography>
              </Box>
              
              <Divider sx={{ my: 2 }} />
              
              <Box>
                <Button 
                  variant="contained" 
                  onClick={() => setEditMode(true)}
                >
                  Edit Profile
                </Button>
              </Box>
            </Stack>
          )}
        </CardContent>
      </Card>
    </Container>
  );
};

export default ProfilePage;