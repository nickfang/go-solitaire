package testutils

import (
	"encoding/json"
	"testing"
)

func AssertNoError(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("Unexpected error during test: %s - %v", message, err)
	}
}

// json
// Helper function to compare JSON strings, ignoring whitespace
func CompareJSON(t *testing.T, expected, actual string) {
	var expectedJSON, actualJSON map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedJSON); err != nil {
		t.Error("Error unmarshalling expected JSON:", err)
	}
	if err := json.Unmarshal([]byte(actual), &actualJSON); err != nil {
		t.Error("Error unmarshalling actual JSON:", err)
	}

	if !EqualJSON(expectedJSON, actualJSON) {
		t.Errorf("JSON mismatch\nExpected: %s\nActual: %s", expected, actual)
	}
}

// Recursive function to compare JSON objects
func EqualJSON(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || !CompareValues(v, w) {
			return false
		}
	}
	return true
}

// Helper function to compare values in JSON objects
func CompareValues(a, b interface{}) bool {
	switch ta := a.(type) {
	case map[string]interface{}:
		tb, ok := b.(map[string]interface{})
		return ok && EqualJSON(ta, tb)
	case []interface{}:
		tb, ok := b.([]interface{})
		return ok && EqualSlices(ta, tb)
	default:
		return a == b
	}
}

// Helper function to compare slices in JSON objects
func EqualSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !CompareValues(v, b[i]) {
			return false
		}
	}
	return true
}
