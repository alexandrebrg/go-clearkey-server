package repositories

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
)

type memory struct {
	keys map[string][]byte
}

func NewMemoryRepository() *memory {
	parse, _ := uuid.Parse("121a0fca-0f1b-475b-8910-297fa8e0a07e")
	key := domain2.ClearKeyDecoded{
		Type:  "temporary",
		Id:    parse,
		Value: uuid.New(),
	}
	bytes, _ := json.Marshal(key)
	return &memory{
		keys: map[string][]byte{
			"121a0fca-0f1b-475b-8910-297fa8e0a07e": bytes,
		},
	}
}

func (repo *memory) Get(id string) (domain2.ClearKeyDecoded, error) {
	if value, ok := repo.keys[id]; ok {
		key := domain2.ClearKeyDecoded{}
		err := json.Unmarshal(value, &key)
		if err != nil {
			logger2.Log.Debugf("GET %s - COULD NOT MARSHAL", id[:8])
			return domain2.ClearKeyDecoded{}, err
		}

		logger2.Log.Debugf("GET %s - FOUND", id[:8])
		return key, nil
	}

	logger2.Log.Debugf("GET %s - NOT FOUND", id[:8])
	return domain2.ClearKeyDecoded{}, errors.New("not found")
}

func (repo *memory) Save(key domain2.ClearKeyDecoded) error {
	bytes, err := json.Marshal(key)

	if err != nil {
		return errors.New("cannot unmarshal")
	}

	repo.keys[key.Id.String()] = bytes
	return nil
}
