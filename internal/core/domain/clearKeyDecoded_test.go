package domain

import (
	"testing"
)

/*
Because if it succeed while ID is nil it won't bother us
*/
func TestClearKeyEncodeSucceedIfNilId(t *testing.T) {
	mockClearKey := MakeClearKeyDecoded().SetNilContent().RandomizeValidValue().Get()

	_, err := mockClearKey.Encode()

	if err != nil {
		t.Errorf("this should not fail as id is nil")
	}
}

func TestClearKeyEncodeSucceedIfNilValue(t *testing.T) {
	mockClearKey := MakeClearKeyDecoded().SetNilContent().RandomizeValidId().Get()

	_, err := mockClearKey.Encode()

	if err != nil {
		t.Errorf("this should not fail as value is nil")
	}
}

func TestClearKeyEncodeSucceedIfValid(t *testing.T) {
	mockClearKey := MakeClearKeyDecoded().RandomizeValidValue().RandomizeValidId().Get()

	_, err := mockClearKey.Encode()

	if err != nil {
		t.Errorf("this should not fail as valid clearkey")
	}
}
