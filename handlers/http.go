package handlers

import (
	"gitlab.com/protocole/clearkey/core/domain"
	"gitlab.com/protocole/clearkey/core/domain/dto"
	"gitlab.com/protocole/clearkey/core/ports/httpHandler"
	"gitlab.com/protocole/clearkey/core/ports/logger"
	"gitlab.com/protocole/clearkey/core/ports/serializer"
	"gitlab.com/protocole/clearkey/core/ports/services"
	"gitlab.com/protocole/clearkey/pkg/apperrors"
	js "gitlab.com/protocole/clearkey/serializers/json"
	"io/ioutil"
	"net/http"
)

type handler struct {
	svc services.KeyStorageService
}

func NewHandler(keyService services.KeyStorageService) httpHandler.HttpHandler {
	return &handler{
		svc: keyService,
	}
}

func (h *handler) setupResponse(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)

	if err != nil {
		logger.Log.Errorf("Failed to response, %s", err.Error())
	}
}

func (h *handler) serializer() serializer.KeyRequestSerializer {
	return &js.RequestKey{}
}

func (h *handler) GetKey(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Errorf("Wasn't able to read from request body, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	keyRequest, err := h.serializer().DecodeRequest(requestBody)
	if err != nil {
		logger.Log.Errorf("Wasn't able to marshal body, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.InvalidInput.WithData(err.Error()))
		return
	}

	keyRequestDecoded, err := keyRequest.Decode()
	if err != nil {
		logger.Log.Errorf("Wasn't able to decode ids, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.InvalidInput.WithData(err.Error()))
		return
	}

	var decodedKeys []domain.ClearKeyEncoded

	for _, keyId := range keyRequestDecoded.KeyIds {
		fetchKey, err := h.svc.GetEncoded(keyId.String())
		if err != nil {
			logger.Log.Debugf("Could not find key (%s)", keyId.String()[:8])
			apperrors.ReturnHttpError(w, apperrors.NotFound.WithData(keyId.String()))
			return
		}

		logger.Log.Debugf("Found key for kid %s", keyId.String())
		decodedKeys = append(decodedKeys, fetchKey)
	}

	response := dto.BuildResponseRequestKey(keyRequestDecoded.SessionType, decodedKeys)

	responseBody, err := h.serializer().EncodeRequest(&response)
	if err != nil {
		logger.Log.Errorf("Could not serialize response: %s", response)
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	h.setupResponse(w, responseBody, http.StatusOK)
}

func (h *handler) PostKey(w http.ResponseWriter, r *http.Request) {
	createdKey, err := h.svc.Create()
	if err != nil {
		logger.Log.Errorf(err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	encodedKey, err := createdKey.Encode()
	if err != nil {
		logger.Log.Errorf("could not encode newly generated key, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	responseBody, err := h.serializer().EncodeKey(&encodedKey)
	if err != nil {
		logger.Log.Errorf("Could not serialize newly generated key, reason: %s", err.Error())
		apperrors.ReturnHttpError(w, apperrors.Internal)
		return
	}

	h.setupResponse(w, responseBody, http.StatusCreated)
}
