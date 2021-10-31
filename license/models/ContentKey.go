package models


type ContentKey struct {
	Type 				string	`json:"kty"`
	ValueAsBase64Url	string 	`json:"k"`
	IdAsBase64Url		string 	`json:"kid"`
}