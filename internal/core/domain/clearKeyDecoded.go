package domain

import (
	"errors"
	"github.com/google/uuid"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
)

type ClearKeyDecoded struct {
	Type 				string
	Id					uuid.UUID
	Value 				uuid.UUID
}

func NewClearKey(cleanId uuid.UUID, cleanValue uuid.UUID, cleanType string) ClearKeyDecoded {
	return ClearKeyDecoded{
		Type:  cleanType,
		Id:    cleanId,
		Value: cleanValue,
	}
}

func (key *ClearKeyDecoded) Encode() (ClearKeyEncoded, error) {
	keyAsB64, err := UUIDToBase64URL(key.Id)
	if err != nil {
		logger2.Log.Errorf("Could not marshal id (%s) into uuid, reason: %s", key.Id, err)
		return ClearKeyEncoded{}, errors.New("could not encode key")
	}

	valueAsB64, err := UUIDToBase64URL(key.Value)
	if err != nil {
		logger2.Log.Errorf("Could not marchal id (%s) into uuid, reason: %s", key.Value, err)
		return ClearKeyEncoded{}, errors.New("could not encode value")
	}
	return ClearKeyEncoded{
		Type:             key.Type,
		ValueAsBase64Url: valueAsB64,
		IdAsBase64Url:    keyAsB64,
	}, nil
}
