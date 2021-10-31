package license

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/helpers"
	"gitlab.com/protocole/clearkey/license/models"
	"gitlab.com/protocole/clearkey/logger"
	"net/http"
)

var Keys = make(map[string]models.ContentKey)

func HandleRequest(c *gin.Context) {
	var request models.LicenseRequest

	if err := c.BindJSON(&request); err != nil {
		logger.Log.Info(err)
		return
	}

	var keys []models.ContentKey

	for keyIndex, keyEncoded := range request.KeyIdsAsBase64Url {
		key, err := helpers.Base64UrlToDecodedString(keyEncoded)
		if err != nil {
			logger.Log.Info(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not parse kids key at index %d", keyIndex),
			})
			return
		}

		keyData, ok := Keys[key]

		if !ok {
			logger.Log.Info("Could not find request key")
			return
		}

		if _, err := uuid.Parse(key); err != nil {
			logger.Log.Info("Could not parse %s to an UUIDv4", key)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not parse kids key at index %d", keyIndex),
			})
			return
		}
		keys = append(keys, models.ContentKey{
			IdAsBase64Url:    keyData.IdAsBase64Url,
			ValueAsBase64Url: keyData.ValueAsBase64Url,
		})
	}

	response := models.LicenseResponse{
		Type: request.SessionType,
		Keys: keys,
	}

	c.IndentedJSON(http.StatusCreated, response)
}
