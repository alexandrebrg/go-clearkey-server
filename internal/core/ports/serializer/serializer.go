package serializer

import (
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	dto2 "gitlab.com/protocole/clearkey/internal/core/domain/dto"
)

type KeyRequestSerializer interface {
	DecodeRequest(input []byte) (*dto2.RequestKeyDto, error)
	EncodeRequest(input *dto2.ResponseRequestKeyDto) ([]byte, error)

	EncodeKey(input *domain2.ClearKeyEncoded) ([]byte, error)
}
