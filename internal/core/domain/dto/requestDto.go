package dto

import (
	"fmt"
	"github.com/google/uuid"
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
)

type RequestKeyDto struct {
	KeyIdsAsBase64Url		[]string 	`json:"kids"`
	SessionType 			string		`json:"type"`
}

type ResponseRequestKeyDto struct {
	Keys	[]domain2.ClearKeyEncoded `json:"keys"`
	Type	string                    `json:"type"`
}

func BuildResponseRequestKey(cleanType string, keys []domain2.ClearKeyEncoded) ResponseRequestKeyDto {
	return ResponseRequestKeyDto{
		Keys: keys,
		Type: cleanType,
	}
}

func (r *RequestKeyDto) Decode() (domain2.RequestKey, error) {
	var decodedKeys []uuid.UUID

	for keyIndex, keyEncoded := range r.KeyIdsAsBase64Url {
		keyAsUUID, err := domain2.Base64URLToUUID(keyEncoded)
		if err != nil {
			logger2.Log.Errorf("Could not decode key (%s), reason: %s", keyEncoded, err)
			return domain2.RequestKey{}, fmt.Errorf("could not decode kid at index %d", keyIndex)
		}

		decodedKeys = append(decodedKeys, keyAsUUID)
	}

	return domain2.RequestKey{
		SessionType: r.SessionType,
		KeyIds: decodedKeys,
	}, nil
}
