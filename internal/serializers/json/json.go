package json

import (
	"encoding/json"
	"errors"
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	dto2 "gitlab.com/protocole/clearkey/internal/core/domain/dto"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
)

type RequestKey struct {}

func (k *RequestKey) DecodeRequest(input []byte) (*dto2.RequestKeyDto, error) {
	keyRequest := &dto2.RequestKeyDto{}
	if err := json.Unmarshal(input, keyRequest); err != nil {
		logger2.Log.Errorf("Decode error %s", err.Error())
		return nil, errors.New("cannot decode right now")
	}

	return keyRequest, nil
}

func (k *RequestKey) EncodeRequest(input *dto2.ResponseRequestKeyDto) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.New("cannot encode right now")
	}

	return rawMsg, nil
}

func (k *RequestKey) EncodeKey(input *domain2.ClearKeyEncoded) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.New("cannot encode right now")
	}

	return rawMsg, nil
}

