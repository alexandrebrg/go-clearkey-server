package serializer

import (
	"gitlab.com/protocole/clearkey/core/domain"
	"gitlab.com/protocole/clearkey/core/domain/dto"
)

type KeyRequestSerializer interface {
	DecodeRequest(input []byte) (*dto.RequestKeyDto, error)
	EncodeRequest(input *dto.ResponseRequestKeyDto) ([]byte, error)

	EncodeKey(input *domain.ClearKeyEncoded) ([]byte, error)
}