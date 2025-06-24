import React from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { CssBaseline, ThemeProvider } from '@mui/material';
import { QueryClient, QueryClientProvider } from 'react-query';

// Theme
import theme from '@/theme/theme';

// Layout
import MainLayout from '@/components/common/Layout/MainLayout';
import AuthLayout from '@/components/common/Layout/AuthLayout';

// Pages
import DashboardPage from '@/pages/Dashboard/DashboardPage';
import CampaignsPage from '@/pages/Campaigns/CampaignsPage';
import CampaignCreatePage from '@/pages/Campaigns/CampaignCreatePage';
import CampaignEditPage from '@/pages/Campaigns/CampaignEditPage';
import DonationsPage from '@/pages/Donations/DonationsPage';
import UsersPage from '@/pages/Users/UsersPage';
import LoginPage from '@/pages/Login/LoginPage';
import NotFoundPage from '@/pages/NotFoundPage';

// Auth Context
import { AuthProvider } from '@/contexts/AuthContext';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <BrowserRouter>
          <AuthProvider>
            <Routes>
              {/* Public Routes */}
              <Route path="/login" element={<AuthLayout><LoginPage /></AuthLayout>} />

              {/* Private Routes */}
              <Route element={<MainLayout />}>
                <Route path="/dashboard" element={<DashboardPage />} />
                <Route path="/campaigns" element={<CampaignsPage />} />
                <Route path="/campaigns/create" element={<CampaignCreatePage />} />
                <Route path="/campaigns/edit/:id" element={<CampaignEditPage />} />
                <Route path="/donations" element={<DonationsPage />} />
                <Route path="/users" element={<UsersPage />} />
                <Route path="/" element={<Navigate to="/dashboard" replace />} />
              </Route>

              {/* 404 */}
              <Route path="*" element={<NotFoundPage />} />
            </Routes>
          </AuthProvider>
        </BrowserRouter>
      </ThemeProvider>
    </QueryClientProvider>
  );
}

export default App;