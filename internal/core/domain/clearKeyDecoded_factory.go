package domain

import (
	"github.com/google/uuid"
)

type ClearKeyDecodedFactory struct {
	item ClearKeyDecoded
}

func MakeClearKeyDecoded() ClearKeyDecodedFactory {
	clearKeyEncoded := ClearKeyDecodedFactory{item: ClearKeyDecoded{
		Type:             "oct",
		Id: 		uuid.New(),
		Value:    uuid.New(),
	}}

	return clearKeyEncoded
}

func (r ClearKeyDecodedFactory) RandomizeValidId() ClearKeyDecodedFactory {
	r.item.Id = uuid.New()
	return r
}

func (r ClearKeyDecodedFactory) RandomizeValidValue() ClearKeyDecodedFactory {
	r.item.Value = uuid.New()
	return r
}

func (r ClearKeyDecodedFactory) SetNilContent() ClearKeyDecodedFactory {
	r.item = ClearKeyDecoded{}
	return r
}

func (r ClearKeyDecodedFactory) Get() ClearKeyDecoded {
	return r.item
}
