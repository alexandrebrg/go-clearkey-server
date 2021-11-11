package license

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/database"
	"gitlab.com/protocole/clearkey/helpers"
	"gitlab.com/protocole/clearkey/license/models"
	"gitlab.com/protocole/clearkey/logger"
	"net/http"
)


func HandleRequest(c *gin.Context) {
	logger.Log.Debug("New request for key received")
	var request models.LicenseRequest

	if err := c.BindJSON(&request); err != nil {
		logger.Log.Errorf("Could not parse request, error: %s", err)
		return
	}

	var keys []models.ContentKeyRequest

	for keyIndex, keyEncoded := range request.KeyIdsAsBase64Url {
		key, err := helpers.Base64UrlToDecodedString(keyEncoded)
		if err != nil {
			logger.Log.Errorf("Could not decode keyEncoded, error: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not parse kids key at index %d", keyIndex),
			})
			return
		}

		if !database.DoesKeyExist(key) {
			logger.Log.Errorf("Could not find key %s", key[:8])
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not find kids key at index %d", keyIndex),
			})
			return
		}

		if _, err := uuid.Parse(key); err != nil {
			logger.Log.Errorf("Could not parse %s to an UUIDv4", key)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not parse kids key at index %d", keyIndex),
			})
			return
		}

		contentKey, err := database.GetKey(key)
		if err != nil {
			return
		}

		base64Value, err := helpers.StringToBase64Url(contentKey.Value)
		if err != nil {
			return
		}

		keys = append(keys, models.ContentKeyRequest{
			IdAsBase64Url:    keyEncoded,
			ValueAsBase64Url: base64Value,
		})
	}

	response := models.LicenseResponse{
		Type: request.SessionType,
		Keys: keys,
	}

	c.IndentedJSON(http.StatusCreated, response)
}
