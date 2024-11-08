package httpresponse_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/zeroxsolutions/go-rps/httpresponse"
	"github.com/zeroxsolutions/go-rps/rpsutil"
)

// TestHTTPResponseBuilder_SimpleSuccess tests the HTTPResponseBuilder with simple success case.
func TestHTTPResponseBuilder_SimpleSuccess(t *testing.T) {
	builder := httpresponse.HTTPResponse[int, string, map[string]interface{}, int]()
	builder.SetMessage("Operation successful").SetCode(200).SetData("Test data").SetTotal(100)

	response, err := rpsutil.Build[httpresponse.HTTPResponseOptions[int, string, map[string]interface{}, int]](builder)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err) // Use t.Fatalf to exit test if error is encountered
	}
	if response == nil {
		t.Fatal("Expected response to be non-nil") // Use t.Fatal to immediately stop if response is nil
	}

	if response.Success != true {
		t.Errorf("Expected Success to be true, got %v", response.Success)
	}
	if response.Message != "Operation successful" {
		t.Errorf("Expected Message to be 'Operation successful', got %v", response.Message)
	}
	if response.Code != 200 {
		t.Errorf("Expected Code to be 200, got %v", response.Code)
	}
	if response.Data != "Test data" {
		t.Errorf("Expected Data to be 'Test data', got %v", response.Data)
	}
	if response.Total != 100 {
		t.Errorf("Expected Total to be 100, got %v", response.Total)
	}
}

// TestHTTPResponseOptions_MarshalJSON tests that custom MarshalJSON includes Extra fields correctly.
func TestHTTPResponseOptions_MarshalJSON(t *testing.T) {
	response := httpresponse.HTTPResponseOptions[int, string, map[string]interface{}, int]{
		Success: true,
		Message: "Success message",
		Code:    200,
		Data:    "Sample data",
		Total:   5,
		Extra: map[string]interface{}{
			"customField1": "value1",
			"customField2": 42,
		},
	}

	jsonData, err := response.MarshalJSON()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify core fields are present in JSON
	jsonString := string(jsonData)
	if !contains(jsonString, `"success":true`) {
		t.Errorf("Expected JSON to contain 'success': true, got %v", jsonString)
	}
	if !contains(jsonString, `"message":"Success message"`) {
		t.Errorf("Expected JSON to contain 'message': 'Success message', got %v", jsonString)
	}
	if !contains(jsonString, `"code":200`) {
		t.Errorf("Expected JSON to contain 'code': 200, got %v", jsonString)
	}
	if !contains(jsonString, `"data":"Sample data"`) {
		t.Errorf("Expected JSON to contain 'data': 'Sample data', got %v", jsonString)
	}
	if !contains(jsonString, `"total":5`) {
		t.Errorf("Expected JSON to contain 'total': 5, got %v", jsonString)
	}

	// Verify Extra fields are included in JSON
	if !contains(jsonString, `"customField1":"value1"`) {
		t.Errorf("Expected JSON to contain 'customField1': 'value1', got %v", jsonString)
	}
	if !contains(jsonString, `"customField2":42`) {
		t.Errorf("Expected JSON to contain 'customField2': 42, got %v", jsonString)
	}
}

// TestHTTPResponseBuilder_ComplexCase tests the builder with more complex configurations.
func TestHTTPResponseBuilder_ComplexCase(t *testing.T) {
	builder := httpresponse.HTTPResponse[int, string, map[string]interface{}, int]()
	builder.SetSuccess(false).
		SetMessage("Complex case error").
		SetCode(500).
		SetData("Error data").
		SetExtra(map[string]interface{}{
			"errorDetail": "Invalid input",
			"retryable":   false,
		}).
		SetTotal(0)

	response, err := rpsutil.Build[httpresponse.HTTPResponseOptions[int, string, map[string]interface{}, int]](builder)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response to be non-nil")
	}

	if response.Success != false {
		t.Errorf("Expected Success to be false, got %v", response.Success)
	}
	if response.Message != "Complex case error" {
		t.Errorf("Expected Message to be 'Complex case error', got %v", response.Message)
	}
	if response.Code != 500 {
		t.Errorf("Expected Code to be 500, got %v", response.Code)
	}
	if response.Data != "Error data" {
		t.Errorf("Expected Data to be 'Error data', got %v", response.Data)
	}
	if response.Extra["errorDetail"] != "Invalid input" || response.Extra["retryable"] != false {
		t.Errorf("Expected Extra to contain 'errorDetail': 'Invalid input' and 'retryable': false, got %v", response.Extra)
	}
	if response.Total != 0 {
		t.Errorf("Expected Total to be 0, got %v", response.Total)
	}
}

// TestBuild_NilOptions tests that Build handles nil options correctly.
func TestBuild_NilOptions(t *testing.T) {
	response, err := rpsutil.Build[httpresponse.HTTPResponseOptions[int, string, map[string]interface{}, int]](httpresponse.HTTPResponse[int, string, map[string]interface{}, int]())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response to be non-nil")
	}

	// By default, the builder should set Success to true
	if response.Success != true {
		t.Errorf("Expected Success to be true by default, got %v", response.Success)
	}
}

// TestBuild_MultipleOptions tests that Build can combine multiple option providers.
func TestBuild_MultipleOptions(t *testing.T) {
	builder1 := httpresponse.HTTPResponse[int, string, map[string]interface{}, int]().
		SetMessage("First builder").
		SetCode(100)

	builder2 := httpresponse.HTTPResponse[int, string, map[string]interface{}, int]().
		SetData("Combined data").
		SetExtra(map[string]interface{}{
			"key1": "value1",
		})

	response, err := rpsutil.Build[httpresponse.HTTPResponseOptions[int, string, map[string]interface{}, int]](builder1, builder2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response to be non-nil")
	}

	if response.Message != "First builder" {
		t.Errorf("Expected Message to be 'First builder', got %v", response.Message)
	}
	if response.Code != 100 {
		t.Errorf("Expected Code to be 100, got %v", response.Code)
	}
	if response.Data != "Combined data" {
		t.Errorf("Expected Data to be 'Combined data', got %v", response.Data)
	}
	if response.Extra["key1"] != "value1" {
		t.Errorf("Expected Extra to contain 'key1': 'value1', got %v", response.Extra)
	}
}

// Helper function to check if a substring is in a string
func contains(str, substr string) bool {
	return json.Valid([]byte(str)) && strings.Contains(str, substr)
}
