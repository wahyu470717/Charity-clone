import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import CssBaseline from '@mui/material/CssBaseline';
import { Box } from '@mui/material';

import { AuthProvider } from './contexts/AuthContext';
import CustomAppBar from './components/common/AppBar';
import ProtectedRoute from './components/common/ProtectedRoute';

// Pages
import HomePage from './pages/home/Home';
import LoginPage from './pages/auth/Login';
import RegisterPage from './pages/auth/Register';
import CampaignListPage from './pages/campaign/CampaignList';
import CampaignDetailPage from './pages/campaign/CampaignDetail';
import CreateCampaignPage from './pages/campaign/CreateCampaign';
import DonatePage from './pages/donation/Donate';
import ProfilePage from './pages/profile/Profile';
import AboutPage from './pages/home/About';
import ContactPage from './pages/home/Contact';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
});

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',
    },
    secondary: {
      main: '#dc004e',
    },
  },
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider theme={theme}>
        <LocalizationProvider dateAdapter={AdapterDateFns}>
          <CssBaseline />
          <AuthProvider>
            <Router>
              <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
                <CustomAppBar />
                <Box component="main" sx={{ flexGrow: 1, pt: 8 }}>
                  <Routes>
                    {/* Redirect root path to login */}
                    <Route path="/" element={<Navigate to="/home" replace />} />
                    
                    {/* Public Routes */}
                    <Route path="/home" element={<HomePage />} />
                    <Route path="/about" element={<AboutPage />} />
                    <Route path="/contact" element={<ContactPage />} />
                    <Route path="/campaigns" element={<CampaignListPage />} />
                    <Route path="/campaigns/:id" element={<CampaignDetailPage />} />
                    
                    {/* Auth Routes */}
                    <Route path="/login" element={<LoginPage />} />
                    <Route path="/register" element={<RegisterPage />} />
                    
                    {/* Protected Routes */}
                    <Route path="/profile" element={
                      <ProtectedRoute>
                        <ProfilePage />
                      </ProtectedRoute>
                    } />
                    <Route path="/campaigns/create" element={
                      <ProtectedRoute roles={['recipient', 'superadmin']}>
                        <CreateCampaignPage />
                      </ProtectedRoute>
                    } />
                    <Route path="/campaigns/:id/donate" element={
                      <ProtectedRoute>
                        <DonatePage />
                      </ProtectedRoute>
                    } />
                  </Routes>
                </Box>
              </Box>
            </Router>
          </AuthProvider>
        </LocalizationProvider>
      </ThemeProvider>
    </QueryClientProvider>
  );
}

export default App;