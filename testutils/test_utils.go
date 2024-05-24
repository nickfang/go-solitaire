package testutils

import (
	"testing"
)

func AssertNoError(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("Unexpected error during test: %s - %v", message, err)
	}
}
