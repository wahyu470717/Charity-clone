import { CampaignCategory, CampaignStatus, DonationStatus, PaymentMethod } from '../types';

// API Configuration
export const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'http://localhost:3001/api';
export const API_VERSION = 'v1';
export const API_TIMEOUT = 30000;

// Authentication
export const TOKEN_KEY = 'charity_admin_token';
export const REFRESH_TOKEN_KEY = 'charity_admin_refresh_token';
export const USER_KEY = 'charity_admin_user';

// Pagination
export const DEFAULT_PAGE_SIZE = 10;
export const PAGE_SIZE_OPTIONS = [5, 10, 25, 50, 100];

// File Upload
export const MAX_FILE_SIZE = 5 * 1024 * 1024; // 5MB
export const ALLOWED_IMAGE_TYPES = ['image/jpeg', 'image/png', 'image/webp'];
export const MAX_IMAGES_PER_CAMPAIGN = 5;

// Campaign Status Options
export const CAMPAIGN_STATUS_OPTIONS: { value: CampaignStatus; label: string; color: string }[] = [
  { value: 'draft', label: 'Draft', color: '#9E9E9E' },
  { value: 'active', label: 'Active', color: '#4CAF50' },
  { value: 'completed', label: 'Completed', color: '#2196F3' },
  { value: 'suspended', label: 'Suspended', color: '#FF9800' },
  { value: 'expired', label: 'Expired', color: '#F44336' },
];

// Donation Status Options
export const DONATION_STATUS_OPTIONS: { value: DonationStatus; label: string; color: string }[] = [
  { value: 'pending', label: 'Pending', color: '#FF9800' },
  { value: 'completed', label: 'Completed', color: '#4CAF50' },
  { value: 'failed', label: 'Failed', color: '#F44336' },
  { value: 'refunded', label: 'Refunded', color: '#9E9E9E' },
];

// Payment Method Options
export const PAYMENT_METHOD_OPTIONS: { value: PaymentMethod; label: string; icon: string }[] = [
  { value: 'credit_card', label: 'Credit Card', icon: 'credit_card' },
  { value: 'bank_transfer', label: 'Bank Transfer', icon: 'account_balance' },
  { value: 'e_wallet', label: 'E-Wallet', icon: 'account_balance_wallet' },
  { value: 'qris', label: 'QRIS', icon: 'qr_code' },
];

// Campaign Categories
export const CAMPAIGN_CATEGORIES: CampaignCategory[] = [
  { id: '1', name: 'Kesehatan', icon: 'local_hospital', color: '#F44336' },
  { id: '2', name: 'Pendidikan', icon: 'school', color: '#2196F3' },
  { id: '3', name: 'Lingkungan', icon: 'eco', color: '#4CAF50' },
  { id: '4', name: 'Sosial', icon: 'group', color: '#FF9800' },
  { id: '5', name: 'Bencana Alam', icon: 'warning', color: '#FF5722' },
  { id: '6', name: 'Ekonomi', icon: 'trending_up', color: '#9C27B0' },
  { id: '7', name: 'Hewan', icon: 'pets', color: '#795548' },
  { id: '8', name: 'Lainnya', icon: 'category', color: '#607D8B' },
];

// User Roles
export const USER_ROLES = [
  { value: 'admin', label: 'Administrator' },
  { value: 'moderator', label: 'Moderator' },
  { value: 'user', label: 'User' },
];

// Chart Colors
export const CHART_COLORS = [
  '#00A651',
  '#FF6B35',
  '#1976D2',
  '#FF9800',
  '#9C27B0',
  '#F44336',
  '#4CAF50',
  '#2196F3',
];

// Date Formats
export const DATE_FORMAT = 'dd/MM/yyyy';
export const DATETIME_FORMAT = 'dd/MM/yyyy HH:mm';
export const TIME_FORMAT = 'HH:mm';

// Currency
export const CURRENCY_SYMBOL = 'Rp';
export const CURRENCY_LOCALE = 'id-ID';

// Menu Items
export const MENU_ITEMS = [
  {
    id: 'dashboard',
    title: 'Dashboard',
    icon: 'dashboard',
    path: '/dashboard',
  },
  {
    id: 'campaigns',
    title: 'Campaigns',
    icon: 'campaign',
    path: '/campaigns',
    children: [
      { id: 'campaigns-list', title: 'All Campaigns', path: '/campaigns' },
      { id: 'campaigns-create', title: 'Create Campaign', path: '/campaigns/create' },
      { id: 'campaigns-categories', title: 'Categories', path: '/campaigns/categories' },
    ],
  },
  {
    id: 'donations',
    title: 'Donations',
    icon: 'volunteer_activism',
    path: '/donations',
    children: [
      { id: 'donations-list', title: 'All Donations', path: '/donations' },
      { id: 'donations-pending', title: 'Pending', path: '/donations?status=pending' },
      { id: 'donations-refunds', title: 'Refunds', path: '/donations/refunds' },
    ],
  },
  {
    id: 'users',
    title: 'Users',
    icon: 'people',
    path: '/users',
    children: [
      { id: 'users-list', title: 'All Users', path: '/users' },
      { id: 'users-verification', title: 'Verification', path: '/users/verification' },
    ],
  },
  {
    id: 'analytics',
    title: 'Analytics',
    icon: 'analytics',
    path: '/analytics',
  },
  {
    id: 'reports',
    title: 'Reports',
    icon: 'assessment',
    path: '/reports',
  },
  {
    id: 'settings',
    title: 'Settings',
    icon: 'settings',
    path: '/settings',
  },
];

// Activity Types
export const ACTIVITY_TYPES = {
  CAMPAIGN_CREATED: 'campaign_created',
  CAMPAIGN_UPDATED: 'campaign_updated',
  CAMPAIGN_DELETED: 'campaign_deleted',
  CAMPAIGN_STATUS_CHANGED: 'campaign_status_changed',
  DONATION_RECEIVED: 'donation_received',
  DONATION_REFUNDED: 'donation_refunded',
  USER_CREATED: 'user_created',
  USER_UPDATED: 'user_updated',
  USER_VERIFIED: 'user_verified',
  USER_SUSPENDED: 'user_suspended',
  ADMIN_LOGIN: 'admin_login',
  ADMIN_LOGOUT: 'admin_logout',
};

// Validation Rules
export const VALIDATION_RULES = {
  EMAIL: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
  PHONE: /^(\+62|62|0)8[1-9][0-9]{6,9}$/,
  PASSWORD_MIN_LENGTH: 8,
  CAMPAIGN_TITLE_MIN_LENGTH: 10,
  CAMPAIGN_TITLE_MAX_LENGTH: 100,
  CAMPAIGN_DESCRIPTION_MIN_LENGTH: 50,
  CAMPAIGN_TARGET_MIN: 100000, // Rp 100,000
  CAMPAIGN_TARGET_MAX: 10000000000, // Rp 10 Billion
};

// Error Messages
export const ERROR_MESSAGES = {
  REQUIRED_FIELD: 'Field ini wajib diisi',
  INVALID_EMAIL: 'Format email tidak valid',
  INVALID_PHONE: 'Format nomor telepon tidak valid',
  PASSWORD_TOO_SHORT: `Password minimal ${VALIDATION_RULES.PASSWORD_MIN_LENGTH} karakter`,
  CAMPAIGN_TITLE_TOO_SHORT: `Judul campaign minimal ${VALIDATION_RULES.CAMPAIGN_TITLE_MIN_LENGTH} karakter`,
  CAMPAIGN_TITLE_TOO_LONG: `Judul campaign maksimal ${VALIDATION_RULES.CAMPAIGN_TITLE_MAX_LENGTH} karakter`,
  CAMPAIGN_DESCRIPTION_TOO_SHORT: `Deskripsi campaign minimal ${VALIDATION_RULES.CAMPAIGN_DESCRIPTION_MIN_LENGTH} karakter`,
  CAMPAIGN_TARGET_TOO_LOW: 'Target donasi terlalu rendah',
  CAMPAIGN_TARGET_TOO_HIGH: 'Target donasi terlalu tinggi',
  FILE_TOO_LARGE: 'Ukuran file terlalu besar',
  INVALID_FILE_TYPE: 'Tipe file tidak didukung',
  NETWORK_ERROR: 'Terjadi kesalahan jaringan',
  UNAUTHORIZED: 'Anda tidak memiliki akses',
  FORBIDDEN: 'Akses ditolak',
  NOT_FOUND: 'Data tidak ditemukan',
  SERVER_ERROR: 'Terjadi kesalahan server',
};

// Success Messages
export const SUCCESS_MESSAGES = {
  CAMPAIGN_CREATED: 'Campaign berhasil dibuat',
  CAMPAIGN_UPDATED: 'Campaign berhasil diperbarui',
  CAMPAIGN_DELETED: 'Campaign berhasil dihapus',
  DONATION_PROCESSED: 'Donasi berhasil diproses',
  DONATION_REFUNDED: 'Donasi berhasil dikembalikan',
  USER_CREATED: 'User berhasil dibuat',
  USER_UPDATED: 'User berhasil diperbarui',
  USER_VERIFIED: 'User berhasil diverifikasi',
  DATA_EXPORTED: 'Data berhasil diekspor',
  SETTINGS_SAVED: 'Pengaturan berhasil disimpan',
};

// Local Storage Keys
export const STORAGE_KEYS = {
  THEME: 'charity_admin_theme',
  LANGUAGE: 'charity_admin_language',
  SIDEBAR_COLLAPSED: 'charity_admin_sidebar_collapsed',
  TABLE_SETTINGS: 'charity_admin_table_settings',
  RECENT_SEARCHES: 'charity_admin_recent_searches',
};

// Export Formats
export const EXPORT_FORMATS = [
  { value: 'csv', label: 'CSV', extension: '.csv' },
  { value: 'excel', label: 'Excel', extension: '.xlsx' },
  { value: 'pdf', label: 'PDF', extension: '.pdf' },
];

// Notification Types
export const NOTIFICATION_TYPES = {
  SUCCESS: 'success',
  ERROR: 'error',
  WARNING: 'warning',
  INFO: 'info',
};

// Refresh Intervals (in milliseconds)
export const REFRESH_INTERVALS = {
  DASHBOARD: 30000, // 30 seconds
  DONATIONS: 60000, // 1 minute
  CAMPAIGNS: 300000, // 5 minutes
  USERS: 600000, // 10 minutes
};