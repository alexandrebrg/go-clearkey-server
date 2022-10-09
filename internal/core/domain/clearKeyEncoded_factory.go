package domain

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type ClearKeyEncodedFactory struct {
	item         ClearKeyEncoded
	decodedId    uuid.UUID
	decodedValue uuid.UUID
}

func MakeClearKeyEncoded() ClearKeyEncodedFactory {
	clearKeyEncoded := ClearKeyEncodedFactory{item: ClearKeyEncoded{
		Type:             "oct",
		ValueAsBase64Url: "",
		IdAsBase64Url:    "",
	}}

	clearKeyEncoded.RandomizeValidId()
	clearKeyEncoded.RandomizeValidValue()

	return clearKeyEncoded
}

func (r ClearKeyEncodedFactory) RandomizeInvalidId() ClearKeyEncodedFactory {
	r.item.IdAsBase64Url = gofakeit.DigitN(18)
	return r
}

func (r ClearKeyEncodedFactory) RandomizeInvalidValue() ClearKeyEncodedFactory {
	r.item.ValueAsBase64Url = gofakeit.DigitN(18)
	return r
}

func (r ClearKeyEncodedFactory) RandomizeValidId() ClearKeyEncodedFactory {
	id, uid := randomizedUUIDInBase64()
	r.item.IdAsBase64Url = id
	r.decodedId = uid
	return r
}

func (r ClearKeyEncodedFactory) RandomizeValidValue() ClearKeyEncodedFactory {
	id, uid := randomizedUUIDInBase64()
	r.item.ValueAsBase64Url = id
	r.decodedValue = uid
	return r
}

func (r ClearKeyEncodedFactory) Get() ClearKeyEncoded {
	return r.item
}

func randomizedUUIDInBase64() (string, uuid.UUID) {
	uid := uuid.New()
	id, _ := UUIDToBase64URL(uid)
	return id, uid
}
