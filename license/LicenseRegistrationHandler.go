package license

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/protocole/clearkey/helpers"
	models2 "gitlab.com/protocole/clearkey/license/models"
	"log"
	"net/http"
)

func HandleKeyRegistration(c *gin.Context) {
	var keyRequest models2.ContentKey

	if err := c.BindJSON(&keyRequest); err != nil {
		log.Println(err)
		return
	}

	keyDecoded, err := helpers.Base64UrlToDecodedString(keyRequest.IdAsBase64Url)
	if err != nil {
		log.Println(err)
		return
	}

	Keys[keyDecoded] = keyRequest
	log.Printf("Registrated new key %s with value %s", keyDecoded, keyRequest.ValueAsBase64Url)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Key registrated",
	})

}
