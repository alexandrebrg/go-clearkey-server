package domain

import (
	"errors"
	"gitlab.com/protocole/clearkey/core/ports/logger"
)

type ClearKeyEncoded struct {
	Type 				string	`json:"kty"`
	ValueAsBase64Url	string 	`json:"k"`
	IdAsBase64Url		string 	`json:"kid"`
}

func (key *ClearKeyEncoded) Decode() (ClearKeyDecoded, error) {
	keyAsUUID, err := Base64URLToUUID(key.IdAsBase64Url)
	if err != nil {
		logger.Log.Errorf("Could not decode key (%s), reason: %s", key.IdAsBase64Url, err)
		return ClearKeyDecoded{}, errors.New("could not decode key")
	}

	valueAsUUID, err := Base64URLToUUID(key.ValueAsBase64Url)
	if err != nil {
		logger.Log.Errorf("Could not decode key (%s), reason: %s", key.ValueAsBase64Url, err)
		return ClearKeyDecoded{}, errors.New("could not decode value")
	}

	return ClearKeyDecoded{
		Type:             key.Type,
		Value: valueAsUUID,
		Id:    keyAsUUID,
	}, nil
}

