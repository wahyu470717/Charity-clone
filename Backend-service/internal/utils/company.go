package utils

import "sync"

type CompanyProfile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Website     string `json:"website"`
}

var (
	companyProfile *CompanyProfile
	profileMutex   sync.RWMutex
)

func init() {
	// Default company profile
	companyProfile = &CompanyProfile{
		Name:        "Share The Meal",
		Description: "Fight hunger worldwide with just a tap",
		Address:     "123 Charity Street, New York, NY",
		Phone:       "+1 (555) 123-4567",
		Email:       "contact@sharethemeal.org",
		Website:     "https://sharethemeal.org",
	}
}

func GetCompanyProfile() *CompanyProfile {
	profileMutex.RLock()
	defer profileMutex.RUnlock()
	return companyProfile
}

func UpdateCompanyProfile(profile CompanyProfile) {
	profileMutex.Lock()
	defer profileMutex.Unlock()
	companyProfile = &profile
}