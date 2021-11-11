package models

type LicenseResponse struct {
	Keys	[]ContentKeyRequest `json:"keys"`
	Type	string       		`json:"type"`
}


