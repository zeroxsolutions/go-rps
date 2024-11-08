package httpresponse_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/zeroxsolutions/go-rps/httpresponse"
)

// TestHTTPResponseDefaultSuccess verifies that the default Success field is true.
func TestHTTPResponseDefaultSuccess(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	options := buildOptions(builder)

	if !options.Success {
		t.Errorf("Expected Success to be true by default, got %v", options.Success)
	}
}

// TestSetSuccess verifies manual setting of the Success field.
func TestSetSuccess(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	builder.SetSuccess(false)
	options := buildOptions(builder)

	if options.Success {
		t.Errorf("Expected Success to be false, got %v", options.Success)
	}
}

// TestSetMessage verifies manual setting of the Message field.
func TestSetMessage(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	message := "Test message"
	builder.SetMessage(message)
	options := buildOptions(builder)

	if options.Message != message {
		t.Errorf("Expected Message to be %v, got %v", message, options.Message)
	}
}

// TestSetCode verifies manual setting of the Code field.
func TestSetCode(t *testing.T) {
	builder := httpresponse.HTTPResponse[int, string, map[string]interface{}, int]()
	code := 200
	builder.SetCode(code)
	options := buildOptions(builder)

	if options.Code != code {
		t.Errorf("Expected Code to be %v, got %v", code, options.Code)
	}
}

// TestSetData verifies manual setting of the Data field.
func TestSetData(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, string, map[string]interface{}, int]()
	data := "Sample Data"
	builder.SetData(data)
	options := buildOptions(builder)

	if options.Data != data {
		t.Errorf("Expected Data to be %v, got %v", data, options.Data)
	}
}

// TestSetTotal verifies manual setting of the Total field.
func TestSetTotal(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	total := 50
	builder.SetTotal(total)
	options := buildOptions(builder)

	if options.Total != total {
		t.Errorf("Expected Total to be %v, got %v", total, options.Total)
	}
}

// TestSetExtra verifies setting Extra with additional fields.
func TestSetExtra(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	extra := map[string]interface{}{
		"extraField1": "extraValue1",
		"extraField2": 42,
	}
	builder.SetExtra(extra)
	options := buildOptions(builder)

	for key, value := range extra {
		if options.Extra[key] != value {
			t.Errorf("Expected Extra[%v] to be %v, got %v", key, value, options.Extra[key])
		}
	}
}

// TestMarshalJSON verifies JSON encoding includes additional fields in Extra.
func TestMarshalJSON(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, string, map[string]interface{}, int]()
	builder.SetSuccess(true)
	builder.SetMessage("Success message")
	builder.SetCode("200")
	builder.SetData("Sample Data")
	builder.SetTotal(100)
	builder.SetExtra(map[string]interface{}{
		"extraField": "extraValue",
	})

	options := buildOptions(builder)

	// Marshal to JSON
	jsonData, err := json.Marshal(options)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Check JSON contains expected fields
	expectedFields := []string{`"success":true`, `"message":"Success message"`, `"code":"200"`, `"data":"Sample Data"`, `"total":100`, `"extraField":"extraValue"`}
	for _, field := range expectedFields {
		if !jsonContains(jsonData, field) {
			t.Errorf("Expected JSON to contain %v, but it did not", field)
		}
	}
}

// TestEmptyExtra verifies behavior when Extra is not set.
func TestEmptyExtra(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, int, map[string]interface{}, int]()
	options := buildOptions(builder)

	if options.Extra != nil {
		t.Errorf("Expected Extra to be nil when not set, got %v", options.Extra)
	}
}

// TestUnsetFields verifies JSON encoding with fields that are not set.
func TestUnsetFields(t *testing.T) {
	builder := httpresponse.HTTPResponse[string, string, map[string]interface{}, int]()
	builder.SetMessage("Only Message Set")

	options := buildOptions(builder)

	jsonData, err := json.Marshal(options)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Check only the message field is set in JSON
	expectedFields := []string{`"message":"Only Message Set"`}
	unexpectedFields := []string{`"code"`, `"data"`, `"total"`, `"extraField"`}
	for _, field := range expectedFields {
		if !jsonContains(jsonData, field) {
			t.Errorf("Expected JSON to contain %v, but it did not", field)
		}
	}
	for _, field := range unexpectedFields {
		if jsonContains(jsonData, field) {
			t.Errorf("Expected JSON not to contain %v, but it did", field)
		}
	}
}

// Helper function to build HTTPResponseOptions from builder
func buildOptions[C int | string, D any, E map[string]interface{}, T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](
	builder *httpresponse.HTTPResponseBuilder[C, D, E, T],
) *httpresponse.HTTPResponseOptions[C, D, E, T] {
	options := &httpresponse.HTTPResponseOptions[C, D, E, T]{}
	for _, opt := range builder.List() {
		opt(options)
	}
	return options
}

// Helper function to check if JSON contains a specific field
func jsonContains(data []byte, field string) bool {
	return string(data) != "" && json.Valid(data) && (string(data) == field || jsonContainsString(string(data), field))
}

// Helper function to check if a string is present in JSON data
func jsonContainsString(jsonStr, field string) bool {
	return jsonStr != "" && field != "" && (jsonStr == field || (len(jsonStr) >= len(field) && strings.Contains(jsonStr, field)))
}
