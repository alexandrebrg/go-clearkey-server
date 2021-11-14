package domain

import (
	"encoding/base64"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"testing"
)

func TestBase64URLToUUIDFailsIfInvalidBase64(t *testing.T) {
	testedString := gofakeit.DigitN(18)

	_, err := Base64URLToUUID(testedString)

	if err == nil {
		t.Errorf("error was expected as invalid base64URL string")
	}
}

func TestBase64URLToUUIDFailsIfInvalidUUID(t *testing.T) {
	testedString := gofakeit.DigitN(18)
	encodedString := base64.RawURLEncoding.EncodeToString([]byte(testedString))

	_, err := Base64URLToUUID(encodedString)

	if err == nil {
		t.Errorf("error was expected as invalid uuid")
	}
}

func TestBase64URLToUUID(t *testing.T) {
	testedUUID := uuid.New()
	testedUUIDBinary, _ := testedUUID.MarshalBinary()
	encodedUUID := base64.RawURLEncoding.EncodeToString(testedUUIDBinary)

	decoded, err := Base64URLToUUID(encodedUUID)

	if err != nil {
		t.Errorf("should be successful as fully valid string")
	}

	if decoded != testedUUID {
		t.Errorf("uuid returned isn't the same as given")
	}
}

func TestUUIDToBase64URL(t *testing.T) {
	testedUUID := uuid.New()
	testedUUIDBinary, _ := testedUUID.MarshalBinary()
	encodedUUID := base64.RawURLEncoding.EncodeToString(testedUUIDBinary)

	encoded, err := UUIDToBase64URL(testedUUID)

	if err != nil {
		t.Errorf("should be successful as fully valid uuid")
	}

	if encoded != encodedUUID {
		t.Errorf("encoded b64URL isn't the same as calculated")
	}
}