package services

import (
	"errors"
	"github.com/google/uuid"
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	repositories2 "gitlab.com/protocole/clearkey/internal/core/ports/repositories"
)

type service struct {
	keyRepository repositories2.KeyStorageRepository
}

func (svc *service) GetAll() (map[string]domain2.ClearKeyDecoded, error) {
	keys, err := svc.keyRepository.GetAll()
	if err != nil {
		return map[string]domain2.ClearKeyDecoded{}, err
	}

	return keys, nil
}

func (svc *service) GetEncoded(keyId string) (domain2.ClearKeyEncoded, error) {
	key, err := svc.Get(keyId)
	if err != nil {
		return domain2.ClearKeyEncoded{}, errors.New("no key found from service")
	}

	return key.Encode()
}

func NewService(keyRepository repositories2.KeyStorageRepository) *service {
	return &service{
		keyRepository,
	}
}

func (svc *service) Get(id string) (domain2.ClearKeyDecoded, error) {
	key, err := svc.keyRepository.Get(id)
	if err != nil {
		return domain2.ClearKeyDecoded{}, errors.New("no key found from service")
	}

	return key, nil
}

func (svc *service) Create() (domain2.ClearKeyDecoded, error) {
	cleanType := "temporary"
	cleanId := uuid.New()
	cleanValue := uuid.New()

	key := domain2.NewClearKey(cleanId, cleanValue, cleanType)

	if err := svc.keyRepository.Save(key); err != nil {
		return domain2.ClearKeyDecoded{}, errors.New("failed to create model")
	}

	return key, nil
}
