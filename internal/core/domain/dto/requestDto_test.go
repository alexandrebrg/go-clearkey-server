package dto

import (
	"gitlab.com/protocole/clearkey/internal/core/domain"
	"gitlab.com/protocole/clearkey/internal/core/ports/logger"
	"gitlab.com/protocole/clearkey/internal/loggers"
	"testing"
)

func TestMain(m *testing.M) {
	logger.SetLogger(loggers.NewZLogger("environment"))
	m.Run()
}

func TestBuildResponseRequestKey(t *testing.T) {
	keys := []domain.ClearKeyEncoded{
		{Type: "oct", ValueAsBase64Url: "test", IdAsBase64Url: "test"},
		{Type: "oct", ValueAsBase64Url: "test2", IdAsBase64Url: "test2"},
	}

	builtResponse := BuildResponseRequestKey("cleanType", keys)

	if builtResponse.Type != "cleanType" {
		t.Errorf("type of key has been changed during building of response")
	}

	if len(builtResponse.Keys) != len(keys) {
		t.Errorf("number of keys returned isn't the same as ones given")
	}

	for index, key := range keys {
		if builtResponse.Keys[index] != key {
			t.Errorf("key has been altered during build response process")
		}
	}
}

func TestDecodeFailsIfInvalidUUID(t *testing.T) {
	keyRequest := MakeRequestKey().AddInvalidKids(5).Get()
	_, errDecoded := keyRequest.Decode()

	if errDecoded == nil {
		t.Errorf("expected error here keyid is invalid")
	}
}

func TestDecodeSucceedIfValidUUID(t *testing.T) {
	mockRequest := MakeRequestKey().AddValidKIds(3).Get()

	returnedRequest, errDecoded := mockRequest.Decode()

	if errDecoded != nil {
		t.Errorf("expected succeed here keyid is valid")
		return
	}

	if len(returnedRequest.KeyIds) != len(mockRequest.KeyIdsAsBase64Url) {
		t.Errorf("number of keys has been altered")
	}

	if returnedRequest.SessionType != mockRequest.SessionType {
		t.Errorf("returned type has been altered")
	}
}

func TestDecodeSucceedIfNoKIds(t *testing.T) {
	mockRequest := MakeRequestKey().Get()

	_, errDecoded := mockRequest.Decode()

	if errDecoded != nil {
		t.Errorf("expected succeed here request is valid")
		return
	}
}
