package license

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/protocole/clearkey/database"
	"gitlab.com/protocole/clearkey/helpers"
	"gitlab.com/protocole/clearkey/license/models"
	"gitlab.com/protocole/clearkey/loggers"
	"net/http"
)

func HandleKeyRegistration(c *gin.Context) {
	loggers.Log.Debug("Received a registration request")
	var keyRequest models.ContentKeyRequest
	var key models.ContentKey

	if err := c.BindJSON(&keyRequest); err != nil {
		loggers.Log.Errorf("Failed to parse JSON, reason %s", err)
		return
	}

	keyDecoded, err := helpers.Base64UrlToDecodedString(keyRequest.IdAsBase64Url)
	if err != nil {
		loggers.Log.Errorf("Failed to decode IdAsBase64Url, reason %s", err)
		return
	}

	key.Id = keyDecoded
	key.Type = keyRequest.Type
	key.Value = "Some random string"

	generatedKey, err := database.TryRegisterKey(key)
	if err != nil {
		return
	}
	loggers.Log.Infof("New key has been registrated ('%s-...-%s')", generatedKey.Id[:8], generatedKey.Id[len(generatedKey.Id) - 8:])

	c.JSON(http.StatusCreated, gin.H{
		"message": "Key registrated",
	})

}
