package repositories

import (
	"gitlab.com/protocole/clearkey/core/domain"
)

type KeyStorageRepository interface {
	Get(keyId string) (domain.ClearKeyDecoded, error)
	Save(keyModel domain.ClearKeyDecoded) error
}
