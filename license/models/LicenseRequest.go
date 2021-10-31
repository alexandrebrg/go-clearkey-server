package models

type LicenseRequest struct {
	KeyIdsAsBase64Url		[]string 	`json:"kids"`
	SessionType 			string		`json:"type"`
}

