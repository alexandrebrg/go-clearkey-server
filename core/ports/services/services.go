package services

import (
	"gitlab.com/protocole/clearkey/core/domain"
)

type KeyStorageService interface {
	Get(keyId string) (domain.ClearKeyDecoded, error)
	GetEncoded(keyId string) (domain.ClearKeyEncoded, error)
	Create() (domain.ClearKeyDecoded, error)
}
