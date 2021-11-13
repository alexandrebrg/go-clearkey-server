package repositories

import (
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
)

type KeyStorageRepository interface {
	Get(keyId string) (domain2.ClearKeyDecoded, error)
	Save(keyModel domain2.ClearKeyDecoded) error
}
