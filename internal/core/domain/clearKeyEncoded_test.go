package domain

import (
	"gitlab.com/protocole/clearkey/internal/core/ports/logger"
	"gitlab.com/protocole/clearkey/internal/loggers"
	"testing"
)

func TestMain(m *testing.M) {
	logger.SetLogger(loggers.NewZLogger("environment"))
	m.Run()
}

func TestDecodeFailsIfIdInvalid(t *testing.T) {
	key := MakeClearKeyEncoded().RandomizeInvalidId().Get()

	_, err := key.Decode()

	if err == nil {
		t.Errorf("expected error as invalid id is given")
	}
}

func TestDecodeFailsIfValueInvalid(t *testing.T) {
	key := MakeClearKeyEncoded().RandomizeInvalidValue().Get()

	_, err := key.Decode()

	if err == nil {
		t.Errorf("expected error as invalid value is given")
	}
}

func TestDecodeSucceedIfValid(t *testing.T) {
	mockKey := MakeClearKeyEncoded().RandomizeValidId().RandomizeValidValue()
	key := mockKey.Get()

	decodedKey, err := key.Decode()

	if err != nil {
		t.Errorf("did not expect error as valid key")
		return
	}

	if decodedKey.Id != mockKey.decodedId {
		t.Errorf("decodedId is not the same as mocked Id")
	}

	if decodedKey.Value != mockKey.decodedValue {
		t.Errorf("decodedValue is not the same as mocked value")
	}

	if decodedKey.Type != key.Type {
		t.Errorf("type of decoded key has been altered")
	}
}
