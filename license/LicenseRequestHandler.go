package license

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/database"
	"gitlab.com/protocole/clearkey/license/models"
	"gitlab.com/protocole/clearkey/loggers"
	"net/http"
)


func HandleRequest(c *gin.Context) {
	loggers.Log.Debug("New request for key received")
	var request models.LicenseRequest

	if err := c.BindJSON(&request); err != nil {
		loggers.Log.Errorf("Could not parse request, error: %s", err)
		return
	}

	var keys []models.ContentKeyRequest

	for keyIndex, keyEncoded := range request.KeyIdsAsBase64Url {
		uuidBytes, err := base64.RawURLEncoding.DecodeString(keyEncoded)
		if err != nil {
			loggers.Log.Errorf("Could not decode keyEncoded, error: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not parse kids key at index %d", keyIndex),
			})
			return
		}
		// We converted decoded base64 bytes into an uuid (16 bytes)
		stringBytes, err := uuid.FromBytes(uuidBytes)
		if err != nil {
			return
		}
		// We convert back the UUID to 16 bytes
		binary, err := stringBytes.MarshalBinary()
		if err != nil {
			return
		}
		// USER FRIENDLY UUID
		loggers.Log.Infof("ORIGIANAL %s", stringBytes.String())
		// USERFRIENDLY UUID
		loggers.Log.Infof("FROM BYTES TO STRING %s", stringBytes)
		// THIS IS WHAT WE WANT FUCK IT
		loggers.Log.Infof("FROM BYTES TO ENCODED STRING %s", base64.RawURLEncoding.EncodeToString(uuidBytes))
		// THIS IS WHAT WE WANT FUCK IT
		loggers.Log.Infof("OUTPUT %s", base64.RawURLEncoding.EncodeToString([]byte(binary)))

		if !database.DoesKeyExist(stringBytes.String()) {
			loggers.Log.Errorf("Could not find key %s", keyEncoded[:8])
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("could not find kids key at index %d", keyIndex),
			})
			return
		}

		loggers.Log.Infof("DECODED STRIN %s", base64.RawURLEncoding.EncodeToString([]byte("121a0fca-0f1b-475b-8910-297fa8e0a07e")))

		keys = append(keys, models.ContentKeyRequest{
			IdAsBase64Url:    keyEncoded,
			ValueAsBase64Url: "EhoPyg8bR1uJECl_qOCgfg",//base64.RawURLEncoding.EncodeToString([]byte("a0a1a2a3a4a5a6a7a8a9aaabacadaeaf")),
		})
	}
	response := models.LicenseResponse{
		Type: request.SessionType,
		Keys: keys,
	}

	c.IndentedJSON(http.StatusOK, response)
}
