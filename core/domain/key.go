package domain

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func Base64URLToUUID(b64 string) (uuid.UUID, error) {
	b64AsBytes, err := base64.RawURLEncoding.DecodeString(b64)
	if err != nil {
		return uuid.New(), err
	}

	b64AsUUID, err := uuid.FromBytes(b64AsBytes)
	if err != nil {
		return uuid.New(), err
	}

	return b64AsUUID, nil
}

func UUIDToBase64URL(id uuid.UUID) (string, error) {
	idAsBinary, err := id.MarshalBinary()

	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(idAsBinary), nil
}
