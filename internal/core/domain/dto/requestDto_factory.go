package dto

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gitlab.com/protocole/clearkey/internal/core/domain"
)

type requestKeyFactory struct {
	item RequestKeyDto
}

func MakeRequestKey() requestKeyFactory {
	return requestKeyFactory{
		item: RequestKeyDto{
			KeyIdsAsBase64Url: []string{},
			SessionType:       "temporary",
		},
	}
}

func (r requestKeyFactory) SetSessionType(sessionType string) requestKeyFactory {
	r.item.SessionType = sessionType
	return r
}

func (r requestKeyFactory) AddInvalidKids(numberOfKids int) requestKeyFactory {
	for i := 0; i < numberOfKids; i++ {
		insertedString := gofakeit.DigitN(uint(gofakeit.Number(2, 18)))
		r.item.KeyIdsAsBase64Url = append(r.item.KeyIdsAsBase64Url, insertedString)
	}

	return r
}

func (r requestKeyFactory) AddValidKIds(numberOfKids int) requestKeyFactory {
	for i := 0; i < numberOfKids; i++ {
		insertString, _ := domain.UUIDToBase64URL(uuid.New())
		r.item.KeyIdsAsBase64Url = append(r.item.KeyIdsAsBase64Url, insertString)
	}

	return r
}

func (r requestKeyFactory) ClearKIds() requestKeyFactory {
	r.item.KeyIdsAsBase64Url = make([]string, 0)
	return r
}

func (r requestKeyFactory) Get() RequestKeyDto {
	return r.item
}
