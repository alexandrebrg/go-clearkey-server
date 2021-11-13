package helpers

import (
	"encoding/base64"
	"errors"
)

func ByteArrayToBase64ByteArray(bytesArray []byte) ([]byte, error) {
	destBytesArray := make([]byte, 4*(len(bytesArray)/3))
	if 0 == len(bytesArray) {
		return nil, errors.New("Invalid bytesArray provided, got length 0")
	}

	base64.StdEncoding.Encode(destBytesArray, bytesArray)

	return destBytesArray, nil
}

func ByteArrayToBase64Url(bytesArray []byte) (string, error) {
	dstByteArray, err := ByteArrayToBase64ByteArray(bytesArray)

	if err != nil {
		return "", nil
	}

	return string(dstByteArray), nil
}

func Base64UrlToByteArray(base64url string) ([]byte, error) {
	if 0 == len(base64url) {
		return nil, errors.New("Base64URL cannot be empty")
	}

	decodedString, err := base64.StdEncoding.DecodeString(base64url)
	if err != nil {
		return nil, err
	}

	return decodedString, nil
}

func Base64UrlToDecodedString(base64url string) (string, error) {
	decodedBytes, err := Base64UrlToByteArray(base64url)
	if err != nil {
		return "", err
	}

	return string(decodedBytes), nil
}

func StringToBase64Url(content string) (string, error) {
	return ByteArrayToBase64Url([]byte(content))
}

func Base64Encode(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

func Base64Decode(message []byte) (b []byte, err error) {
	var l int
	b = make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, err = base64.StdEncoding.Decode(b, message)
	if err != nil {
		return
	}
	return b[:l], nil
}
