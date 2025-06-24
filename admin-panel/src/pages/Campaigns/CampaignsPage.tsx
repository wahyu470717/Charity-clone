import React, { useState } from 'react';
import { Box, Button, Container, Typography } from '@mui/material';
import { useQuery } from 'react-query';
import { getCampaigns } from '@/services/campaignService';
import DataTable from '@/components/common/DataTable/DataTable';
import { Campaign, CampaignStatus } from '@/types';
import CampaignStatusBadge from '@/components/campaigns/CampaignStatusBadge';
import { Link } from 'react-router-dom';
import { CAMPAIGN_STATUS_OPTIONS } from '@/constants';

const CampaignsPage: React.FC = () => {
  const [filters, setFilters] = useState<Record<string, any>>({});
  const [pagination, setPagination] = useState({ page: 1, limit: 10 });
  
  const { data, isLoading } = useQuery(
    ['campaigns', pagination, filters],
    () => getCampaigns({ ...pagination, ...filters })
  );

  const columns = [
    { id: 'title', label: 'Judul', minWidth: 200 },
    { 
      id: 'status', 
      label: 'Status', 
      minWidth: 100,
      format: (value: CampaignStatus) => (
        <CampaignStatusBadge status={value} />
      )
    },
    { id: 'targetAmount', label: 'Target', minWidth: 100, format: (value: number) => `Rp ${value.toLocaleString('id-ID')}` },
    { id: 'currentAmount', label: 'Terkumpul', minWidth: 100, format: (value: number) => `Rp ${value.toLocaleString('id-ID')}` },
    { id: 'donorCount', label: 'Donatur', minWidth: 50 },
    { id: 'createdAt', label: 'Dibuat', minWidth: 100 },
  ];

  return (
    <Container>
      <Box display="flex" justifyContent="space-between" mb={4}>
        <Typography variant="h5">Kelola Campaign</Typography>
        <Button 
          variant="contained" 
          color="primary" 
          component={Link}
          to="/campaigns/create"
        >
          Buat Campaign Baru
        </Button>
      </Box>

      {/* Filter Component Here */}

      <DataTable 
        columns={columns}
        rows={data?.data || []}
        loading={isLoading}
        rowCount={data?.pagination?.total || 0}
        page={pagination.page - 1}
        onPageChange={(newPage) => setPagination({...pagination, page: newPage + 1})}
        rowsPerPage={pagination.limit}
        onRowsPerPageChange={(limit) => setPagination({...pagination, limit})}
      />
    </Container>
  );
};

export default CampaignsPage;