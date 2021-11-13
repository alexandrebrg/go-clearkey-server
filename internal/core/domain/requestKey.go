package domain

import "github.com/google/uuid"

type RequestKey struct {
	KeyIds			[]uuid.UUID
	SessionType		string
}