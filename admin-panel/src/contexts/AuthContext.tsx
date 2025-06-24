import React, { createContext, useContext, useReducer, useEffect, ReactNode } from 'react';
import { AuthUser, LoginCredentials } from '../types';
import { TOKEN_KEY, REFRESH_TOKEN_KEY, USER_KEY } from '../constants';
import { getLocalStorage, setLocalStorage, removeLocalStorage } from '../utils';
import { authService } from '../services/authService';

interface AuthState {
  user: AuthUser | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

type AuthAction =
  | { type: 'AUTH_START' }
  | { type: 'AUTH_SUCCESS'; payload: { user: AuthUser; token: string } }
  | { type: 'AUTH_ERROR'; payload: string }
  | { type: 'LOGOUT' }
  | { type: 'CLEAR_ERROR' }
  | { type: 'UPDATE_USER'; payload: AuthUser };

interface AuthContextType extends AuthState {
  login: (credentials: LoginCredentials) => Promise<void>;
  logout: () => void;
  refreshAuth: () => Promise<void>;
  updateUser: (user: AuthUser) => void;
  clearError: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const initialState: AuthState = {
  user: null,
  token: null,
  isAuthenticated: false,
  isLoading: true,
  error: null,
};

const authReducer = (state: AuthState, action: AuthAction): AuthState => {
  switch (action.type) {
    case 'AUTH_START':
      return {
        ...state,
        isLoading: true,
        error: null,
      };
    case 'AUTH_SUCCESS':
      return {
        ...state,
        user: action.payload.user,
        token: action.payload.token,
        isAuthenticated: true,
        isLoading: false,
        error: null,
      };
    case 'AUTH_ERROR':
      return {
        ...state,
        user: null,
        token: null,
        isAuthenticated: false,
        isLoading: false,
        error: action.payload,
      };
    case 'LOGOUT':
      return {
        ...state,
        user: null,
        token: null,
        isAuthenticated: false,
        isLoading: false,
        error: null,
      };
    case 'UPDATE_USER':
      return {
        ...state,
        user: action.payload,
      };
    case 'CLEAR_ERROR':
      return {
        ...state,
        error: null,
      };
    default:
      return state;
  }
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [state, dispatch] = useReducer(authReducer, initialState);

  // Initialize auth state from localStorage
  useEffect(() => {
    const initializeAuth = async () => {
      const storedToken = getLocalStorage(TOKEN_KEY, '');
      const storedUser = getLocalStorage<AuthUser | null>(USER_KEY, null);

      if (storedToken && storedUser) {
        try {
          // Verify token is still valid
          authService.setAuthToken(storedToken);
          const response = await authService.verifyToken();
          
          dispatch({
            type: 'AUTH_SUCCESS',
            payload: {
              user: response.user,
              token: storedToken,
            },
          });
        } catch (error) {
          // Token is invalid, clear stored data
          clearStoredAuth();
          dispatch({ type: 'AUTH_ERROR', payload: 'Session expired' });
        }
      } else {
        dispatch({ type: 'AUTH_ERROR', payload: 'Not authenticated' });
      }
    };

    initializeAuth();
  }, []);

  const clearStoredAuth = () => {
    removeLocalStorage(TOKEN_KEY);
    removeLocalStorage(REFRESH_TOKEN_KEY);
    removeLocalStorage(USER_KEY);
    authService.setAuthToken('');
  };

  const login = async (credentials: LoginCredentials) => {
    dispatch({ type: 'AUTH_START' });

    try {
      const response = await authService.login(credentials);
      
      // Store auth data
      setLocalStorage(TOKEN_KEY, response.token);
      setLocalStorage(REFRESH_TOKEN_KEY, response.refreshToken);
      setLocalStorage(USER_KEY, response.user);
      
      // Set auth token for future requests
      authService.setAuthToken(response.token);

      dispatch({
        type: 'AUTH_SUCCESS',
        payload: {
          user: response.user,
          token: response.token,
        },
      });
    } catch (error: any) {
      const errorMessage = error.response?.data?.message || 'Login failed';
      dispatch({ type: 'AUTH_ERROR', payload: errorMessage });
      throw error;
    }
  };

  const logout = async () => {
    try {
      await authService.logout();
    } catch (error) {
      console.error('Logout error:', error);
    } finally {
      clearStoredAuth();
      dispatch({ type: 'LOGOUT' });
    }
  };

  const refreshAuth = async () => {
    const refreshToken = getLocalStorage(REFRESH_TOKEN_KEY, '');
    
    if (!refreshToken) {
      dispatch({ type: 'AUTH_ERROR', payload: 'No refresh token' });
      return;
    }

    try {
      const response = await authService.refreshToken(refreshToken);
      
      // Update stored tokens
      setLocalStorage(TOKEN_KEY, response.token);
      setLocalStorage(REFRESH_TOKEN_KEY, response.refreshToken);
      setLocalStorage(USER_KEY, response.user);
      
      // Set new auth token
      authService.setAuthToken(response.token);

      dispatch({
        type: 'AUTH_SUCCESS',
        payload: {
          user: response.user,
          token: response.token,
        },
      });
    } catch (error: any) {
      clearStoredAuth();
      dispatch({ type: 'AUTH_ERROR', payload: 'Failed to refresh token' });
      throw error;
    }
  };

  const updateUser = (user: AuthUser) => {
    setLocalStorage(USER_KEY, user);
    dispatch({ type: 'UPDATE_USER', payload: user });
  };

  const clearError = () => {
    dispatch({ type: 'CLEAR_ERROR' });
  };

  const value: AuthContextType = {
    ...state,
    login,
    logout,
    refreshAuth,
    updateUser,
    clearError,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};