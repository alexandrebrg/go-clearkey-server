package handlers

import (
	jsonRaw "encoding/json"
	domain2 "gitlab.com/protocole/clearkey/internal/core/domain"
	dto2 "gitlab.com/protocole/clearkey/internal/core/domain/dto"
	httpHandler2 "gitlab.com/protocole/clearkey/internal/core/ports/httpHandler"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
	serializer2 "gitlab.com/protocole/clearkey/internal/core/ports/serializer"
	services2 "gitlab.com/protocole/clearkey/internal/core/ports/services"
	"gitlab.com/protocole/clearkey/internal/serializers/json"
	"gitlab.com/protocole/clearkey/pkg/apperrors"
	"io/ioutil"
	"net/http"
)

type handler struct {
	svc services2.KeyStorageService
}

func NewHandler(keyService services2.KeyStorageService) httpHandler2.HttpHandler {
	return &handler{
		svc: keyService,
	}
}

func (h *handler) setupResponse(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)

	if err != nil {
		logger2.Log.Errorf("Failed to response, %s", err.Error())
	}
}

func (h *handler) serializer() serializer2.KeyRequestSerializer {
	return &json.RequestKey{}
}

func (h *handler) GetKeys(w http.ResponseWriter, r *http.Request) {
	keys, err := h.svc.GetAll()
	if err != nil {
		logger2.Log.Errorf("Could not fetch keys right now")
		apperrors.ReturnHttpError(w, apperrors.Internal)
	}

	responseBody, err := jsonRaw.Marshal(keys)
	if err != nil {
		logger2.Log.Errorf("Could not unmarshal content, reason %s", err)
		apperrors.ReturnHttpError(w, apperrors.Internal)
	}

	h.setupResponse(w, responseBody, http.StatusOK)
}

func (h *handler) GetKey(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger2.Log.Errorf("Wasn't able to read from request body, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	keyRequest, err := h.serializer().DecodeRequest(requestBody)
	if err != nil {
		logger2.Log.Errorf("Wasn't able to marshal body, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.InvalidInput.WithData(err.Error()))
		return
	}

	keyRequestDecoded, err := keyRequest.Decode()
	if err != nil {
		logger2.Log.Errorf("Wasn't able to decode ids, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.InvalidInput.WithData(err.Error()))
		return
	}

	var decodedKeys []domain2.ClearKeyEncoded

	for _, keyId := range keyRequestDecoded.KeyIds {
		fetchKey, err := h.svc.GetEncoded(keyId.String())
		if err != nil {
			logger2.Log.Debugf("Could not find key (%s)", keyId.String()[:8])
			apperrors.ReturnHttpError(w, apperrors.NotFound.WithData(keyId.String()))
			return
		}

		logger2.Log.Debugf("Found key for kid %s", keyId.String())
		decodedKeys = append(decodedKeys, fetchKey)
	}

	response := dto2.BuildResponseRequestKey(keyRequestDecoded.SessionType, decodedKeys)

	responseBody, err := h.serializer().EncodeRequest(&response)
	if err != nil {
		logger2.Log.Errorf("Could not serialize response: %s", response)
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	h.setupResponse(w, responseBody, http.StatusOK)
}

func (h *handler) PostKey(w http.ResponseWriter, r *http.Request) {
	createdKey, err := h.svc.Create()
	if err != nil {
		logger2.Log.Errorf(err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	encodedKey, err := createdKey.Encode()
	if err != nil {
		logger2.Log.Errorf("could not encode newly generated key, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	responseBody, err := h.serializer().EncodeKey(&encodedKey)
	if err != nil {
		logger2.Log.Errorf("Could not serialize newly generated key, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	h.setupResponse(w, responseBody, http.StatusCreated)
}
