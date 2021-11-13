package json

import (
	"encoding/json"
	"errors"
	"gitlab.com/protocole/clearkey/core/domain"
	"gitlab.com/protocole/clearkey/core/domain/dto"
	"gitlab.com/protocole/clearkey/core/ports/logger"
)

type RequestKey struct {}

func (k *RequestKey) DecodeRequest(input []byte) (*dto.RequestKeyDto, error) {
	keyRequest := &dto.RequestKeyDto{}
	if err := json.Unmarshal(input, keyRequest); err != nil {
		logger.Log.Errorf("Decode error %s", err.Error())
		return nil, errors.New("cannot decode right now")
	}

	return keyRequest, nil
}

func (k *RequestKey) EncodeRequest(input *dto.ResponseRequestKeyDto) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.New("cannot encode right now")
	}

	return rawMsg, nil
}

func (k *RequestKey) EncodeKey(input *domain.ClearKeyEncoded) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.New("cannot encode right now")
	}

	return rawMsg, nil
}

