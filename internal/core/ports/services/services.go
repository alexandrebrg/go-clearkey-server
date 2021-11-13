package services

import (
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
)

type KeyStorageService interface {
	Get(keyId string) (domain2.ClearKeyDecoded, error)
	GetEncoded(keyId string) (domain2.ClearKeyEncoded, error)
	Create() (domain2.ClearKeyDecoded, error)
}
