package models

type LicenseResponse struct {
	Keys	[]ContentKey `json:"keys"`
	Type	string       `json:"type"`
}


