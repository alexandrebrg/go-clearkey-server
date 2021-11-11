package models


type ContentKeyRequest struct {
	Type 				string	`json:"kty"`
	ValueAsBase64Url	string 	`json:"k"`
	IdAsBase64Url		string 	`json:"kid"`
}