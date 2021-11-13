package dto

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/core/domain"
	"gitlab.com/protocole/clearkey/core/ports/logger"
)

type RequestKeyDto struct {
	KeyIdsAsBase64Url		[]string 	`json:"kids"`
	SessionType 			string		`json:"type"`
}

type ResponseRequestKeyDto struct {
	Keys	[]domain.ClearKeyEncoded `json:"keys"`
	Type	string                   `json:"type"`
}

func BuildResponseRequestKey(cleanType string, keys []domain.ClearKeyEncoded) ResponseRequestKeyDto {
	return ResponseRequestKeyDto{
		Keys: keys,
		Type: cleanType,
	}
}

func (r *RequestKeyDto) Decode() (domain.RequestKey, error) {
	var decodedKeys []uuid.UUID

	for keyIndex, keyEncoded := range r.KeyIdsAsBase64Url {
		keyAsUUID, err := domain.Base64URLToUUID(keyEncoded)
		if err != nil {
			logger.Log.Errorf("Could not decode key (%s), reason: %s", keyEncoded, err)
			return domain.RequestKey{}, fmt.Errorf("could not decode kid at index %d", keyIndex)
		}

		decodedKeys = append(decodedKeys, keyAsUUID)
	}

	return domain.RequestKey{
		SessionType: r.SessionType,
		KeyIds: decodedKeys,
	}, nil
}
