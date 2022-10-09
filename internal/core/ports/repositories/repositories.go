package repositories

import (
	"gitlab.com/protocole/clearkey/internal/core/domain"
)

type KeyStorageRepository interface {
	Get(keyId string) (domain.ClearKeyDecoded, error)
	GetAll() (map[string]domain.ClearKeyDecoded, error)
	Save(keyModel domain.ClearKeyDecoded) error
}
