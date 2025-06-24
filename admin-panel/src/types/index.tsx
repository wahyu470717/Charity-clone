// User Types
export interface User {
  id: string;
  name: string;
  email: string;
  phone?: string;
  avatar?: string;
  role: 'admin' | 'moderator' | 'user';
  isVerified: boolean;
  createdAt: string;
  lastLogin?: string;
  totalDonations: number;
  totalCampaigns: number;
}

// Campaign Types
export interface Campaign {
  id: string;
  title: string;
  description: string;
  shortDescription: string;
  image: string;
  images: string[];
  targetAmount: number;
  currentAmount: number;
  donorCount: number;
  category: CampaignCategory;
  status: CampaignStatus;
  createdBy: User;
  createdAt: string;
  updatedAt: string;
  endDate: string;
  location: string;
  isUrgent: boolean;
  tags: string[];
}

export type CampaignStatus = 'draft' | 'active' | 'completed' | 'suspended' | 'expired';

export interface CampaignCategory {
  id: string;
  name: string;
  icon: string;
  color: string;
}

// Donation Types
export interface Donation {
  id: string;
  amount: number;
  campaignId: string;
  campaign: Campaign;
  donorId?: string;
  donor?: User;
  donorName: string;
  donorEmail?: string;
  message?: string;
  isAnonymous: boolean;
  status: DonationStatus;
  paymentMethod: PaymentMethod;
  transactionId: string;
  createdAt: string;
  processedAt?: string;
}

export type DonationStatus = 'pending' | 'completed' | 'failed' | 'refunded';
export type PaymentMethod = 'credit_card' | 'bank_transfer' | 'e_wallet' | 'qris';

// Analytics Types
export interface DashboardStats {
  totalCampaigns: number;
  activeCampaigns: number;
  totalDonations: number;
  totalAmount: number;
  totalUsers: number;
  verifiedUsers: number;
  monthlyGrowth: {
    campaigns: number;
    donations: number;
    users: number;
    amount: number;
  };
}

export interface ChartData {
  date: string;
  donations: number;
  amount: number;
  campaigns: number;
}

export interface CategoryStats {
  category: string;
  count: number;
  amount: number;
  percentage: number;
}

// API Response Types
export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
  pagination?: {
    page: number;
    limit: number;
    total: number;
    totalPages: number;
  };
}

export interface PaginationParams {
  page: number;
  limit: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
  search?: string;
}

// Form Types
export interface CampaignFormData {
  title: string;
  description: string;
  shortDescription: string;
  targetAmount: number;
  categoryId: string;
  endDate: string;
  location: string;
  isUrgent: boolean;
  tags: string[];
  images: File[];
}

export interface UserFormData {
  name: string;
  email: string;
  phone?: string;
  role: 'admin' | 'moderator' | 'user';
}

// Auth Types
export interface AuthUser {
  id: string;
  name: string;
  email: string;
  role: 'admin' | 'moderator';
  avatar?: string;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface AuthResponse {
  user: AuthUser;
  token: string;
  refreshToken: string;
}

// Filter Types
export interface CampaignFilters {
  status?: CampaignStatus[];
  category?: string[];
  dateRange?: {
    start: string;
    end: string;
  };
  amountRange?: {
    min: number;
    max: number;
  };
  search?: string;
}

export interface DonationFilters {
  status?: DonationStatus[];
  paymentMethod?: PaymentMethod[];
  dateRange?: {
    start: string;
    end: string;
  };
  amountRange?: {
    min: number;
    max: number;
  };
  campaignId?: string;
  search?: string;
}

// Table Types
export interface TableColumn<T> {
  id: keyof T;
  label: string;
  minWidth?: number;
  align?: 'right' | 'left' | 'center';
  format?: (value: any) => string;
  sortable?: boolean;
}

// Notification Types
export interface Notification {
  id: string;
  type: 'success' | 'error' | 'warning' | 'info';
  title: string;
  message: string;
  createdAt: string;
  isRead: boolean;
}

// Export/Import Types
export interface ExportOptions {
  format: 'csv' | 'excel' | 'pdf';
  dateRange?: {
    start: string;
    end: string;
  };
  columns: string[];
}

// Activity Log Types
export interface ActivityLog {
  id: string;
  userId: string;
  user: User;
  action: string;
  resource: string;
  resourceId: string;
  details: Record<string, any>;
  ipAddress: string;
  userAgent: string;
  createdAt: string;
}