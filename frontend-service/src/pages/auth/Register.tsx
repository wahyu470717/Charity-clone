import React from 'react';
import { Container, Box, Typography, Paper } from '@mui/material';
import RegisterForm from '../../components/auth/RegisterForm';
import { useAuth } from '../../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';
import { authApi } from '../../api/auth';

const RegisterPage: React.FC = () => {
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleRegister = async (userData: {
    name: string;
    email: string;
    password: string;
    role: string;
  }) => {
    await authApi.register(userData);
    // After successful registration, automatically log in
    const { token, user } = await authApi.login(userData.email, userData.password);
    login(user, token);
    navigate('/');
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 8, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
          <Typography component="h1" variant="h4" align="center" gutterBottom>
            Sign up
          </Typography>
          <RegisterForm onSubmit={handleRegister} />
        </Paper>
      </Box>
    </Container>
  );
};

export default RegisterPage;