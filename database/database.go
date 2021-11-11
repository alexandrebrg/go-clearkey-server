package database

import (
	"errors"
	"gitlab.com/protocole/clearkey/license/models"
)

var keys = make(map[string]models.ContentKey)

func TryRegisterKey(key models.ContentKey) (*models.ContentKey, error) {
	if DoesKeyExist(key.Id) {
		return nil, errors.New("key already exists")
	}

	registerKey(key)

	keyRef, err := GetKey(key.Id)
	if err != nil {
		return nil, err
	}

	return keyRef, nil
}

func DoesKeyExist(key string) bool {
	_, ok := keys[key]

	return ok
}

func registerKey(key models.ContentKey) {
	keys[key.Id] = key
}

func GetKey(keyId string) (*models.ContentKey, error) {
	key, ok := keys[keyId]
	if !ok {
		return nil, errors.New("cannot fetch key as it does not exist")
	}

	return &key, nil
}
