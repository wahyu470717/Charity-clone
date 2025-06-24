import React from 'react';
import { Container, Box, Typography, Paper } from '@mui/material';
import LoginForm from '../../components/auth/LoginForm';
import { useAuth } from '../../contexts/AuthContext';
import { useLocation, useNavigate } from 'react-router-dom';
import { authApi } from '../../api/auth';

const LoginPage: React.FC = () => {
  const { login } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogin = async (email: string, password: string) => {
    const { token, user } = await authApi.login(email, password);
    login(user, token);
    
    // Redirect ke halaman sebelumnya atau home
    const from = location.state?.from?.pathname || '/home';
    navigate(from, { replace: true });
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 8, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
          <Typography component="h1" variant="h4" align="center" gutterBottom>
            Sign in
          </Typography>
          <LoginForm onSubmit={handleLogin} />
        </Paper>
      </Box>
    </Container>
  );
};

export default LoginPage;