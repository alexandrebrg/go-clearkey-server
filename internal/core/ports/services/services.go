package services

import (
	"gitlab.com/protocole/clearkey/internal/core/domain"
)

type KeyStorageService interface {
	Get(keyId string) (domain.ClearKeyDecoded, error)
	GetEncoded(keyId string) (domain.ClearKeyEncoded, error)
	Create() (domain.ClearKeyDecoded, error)
	GetAll() (map[string]domain.ClearKeyDecoded, error)
}
