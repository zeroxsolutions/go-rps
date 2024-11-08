// Package httpresponse provides structures and methods to configure HTTP response options
// in a flexible and reusable way. It supports customizable fields for success indicators,
// messages, status codes, data payloads, total counts, and additional metadata.
package httpresponse

import (
	"encoding/json"
)

// HTTPResponseOptions represents the configuration of an HTTP response, with fields that
// can be customized to suit various response needs, such as success status, messages, response
// codes, data payloads, totals, and additional metadata.
//
// Type parameters:
//   - C: Type for the response code, which can be an integer or a string.
//   - D: Type for the data payload, supporting any type (e.g., string, struct, or slice).
//   - E: Type for additional metadata, defined as a map with string keys and any values.
//   - T: Type for the total field, allowing various integer types (e.g., int, uint, int64).
type HTTPResponseOptions[
	C int | string,
	D any,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
] struct {
	Success bool   `json:"success"`         // Indicates if the response signifies a successful operation.
	Message string `json:"message"`         // Descriptive message for the response, such as success or error info.
	Code    C      `json:"code,omitempty"`  // Status code for the response (e.g., HTTP code or custom code); omitted if empty.
	Data    D      `json:"data,omitempty"`  // Payload containing the main response data; omitted if empty.
	Total   T      `json:"total,omitempty"` // Total count or amount, often used for pagination; omitted if empty.
	Extra   E      `json:"-"`               // Additional metadata excluded from JSON by default.
}

// MarshalJSON customizes the JSON encoding for HTTPResponseOptions by merging the core
// fields with any additional metadata provided in the Extra map.
//
// This method first marshals the standard fields of HTTPResponseOptions into JSON, then adds
// any fields from the Extra map into the resulting JSON object before finalizing the output.
//
// Returns:
//   - []byte: The customized JSON encoding of HTTPResponseOptions, with merged Extra fields.
//   - error: An error if the marshaling or merging process fails.
func (httpResponseOptions *HTTPResponseOptions[C, D, E, T]) MarshalJSON() ([]byte, error) {

	// Marshal the core fields into JSON
	r, err := json.Marshal(HTTPResponseOptions[C, D, E, T]{
		Success: httpResponseOptions.Success,
		Message: httpResponseOptions.Message,
		Code:    httpResponseOptions.Code,
		Data:    httpResponseOptions.Data,
		Total:   httpResponseOptions.Total,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal the core fields into a map for merging with Extra fields
	var rm map[string]interface{}
	if err := json.Unmarshal(r, &rm); err != nil {
		return nil, err
	}

	// Integrate Extra fields into the map if they exist
	if httpResponseOptions.Extra != nil {
		for k, v := range httpResponseOptions.Extra {
			rm[k] = v
		}
	}

	// Marshal the combined map (core fields + Extra fields) back to JSON
	return json.Marshal(rm)
}
