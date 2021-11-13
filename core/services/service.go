package services

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/core/domain"
	"gitlab.com/protocole/clearkey/core/ports/repositories"
)

type service struct {
	keyRepository repositories.KeyStorageRepository
}

func (svc *service) GetEncoded(keyId string) (domain.ClearKeyEncoded, error) {
	key, err := svc.Get(keyId)
	if err != nil {
		return domain.ClearKeyEncoded{}, errors.New("no key found from service")
	}

	return key.Encode()
}

func NewService(keyRepository repositories.KeyStorageRepository) *service {
	return &service{
		keyRepository,
	}
}

func (svc *service) Get(id string) (domain.ClearKeyDecoded, error) {
	key, err := svc.keyRepository.Get(id)
	if err != nil {
		return domain.ClearKeyDecoded{}, errors.New("no key found from service")
	}

	return key, nil
}

func (svc *service) Create() (domain.ClearKeyDecoded, error) {
	cleanType := "temporary"
	cleanId := uuid.New()
	cleanValue := uuid.New()

	key := domain.NewClearKey(cleanId, cleanValue, cleanType)

	if err := svc.keyRepository.Save(key); err != nil {
		return domain.ClearKeyDecoded{}, errors.New("failed to create model")
	}

	return key, nil
}
