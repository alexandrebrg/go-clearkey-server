package helpers

import (
	"github.com/google/uuid"
	"testing"
)

func TestConvertAndFindBackUUID(t *testing.T) {
	if testing.Short() {
		t.Skip("Not useful in short mode")
	}
	// Generate a new UUID
	myUUID := uuid.New()
	// EncodeRequest it into base64
	encodedBase64, err := ByteArrayToBase64ByteArray([]byte(myUUID.String()))
	if err != nil {
		t.Errorf("Converting UUID to Base64ByteArray failed %s", err)
	}

	// DecodeRequest the encoded base64
	decodedBase64, err := Base64UrlToByteArray(string(encodedBase64))
	if err != nil {
		t.Errorf("Converting Bas6e4 to ByteArray failed %s", err)
	}
	// Create a new UUID object based on decoded base64
	_, err = uuid.ParseBytes(decodedBase64)
	if err != nil {
		t.Errorf("Converting Bytes to UUID failed %s", err)
	}
}
