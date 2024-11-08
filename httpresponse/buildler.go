// Package httpresponse provides a builder-pattern implementation for constructing flexible and customizable HTTP responses.
// The builder simplifies setting fields such as success status, message, response codes, data, additional metadata,
// and total count, promoting reusable and structured HTTP response creation.
package httpresponse

// HTTPResponseBuilder is a generic builder that facilitates the creation of structured HTTP response configurations.
// It enables the setting of various response fields like success status, message, code, data, total, and extra metadata.
//
// Type parameters:
//   - C: Defines the type for the response code, which can be either an int or a string.
//   - D: Defines the type for the data field, which can be any data type (e.g., string, struct, array, etc.).
//   - E: Defines the type for additional metadata as a map with string keys and any values.
//   - T: Defines the type for the total field, which can be any integer type (e.g., int, uint, int64).
type HTTPResponseBuilder[
	C int | string,
	D any,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
] struct {
	Opts []func(*HTTPResponseOptions[C, D, E, T]) error
}

// HTTPResponse creates a new instance of HTTPResponseBuilder with default settings.
// By default, the Success field is set to true, indicating a successful response.
//
// Returns:
//   - *HTTPResponseBuilder: A new instance of HTTPResponseBuilder with default success status.
func HTTPResponse[
	D any,
	C int | string,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
]() *HTTPResponseBuilder[C, D, E, T] {

	httpResponseBuilder := new(HTTPResponseBuilder[C, D, E, T])

	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Success = true

		return nil
	})

	return httpResponseBuilder
}

// SetSuccess sets the Success field in the HTTP response options.
//
// Parameters:
//   - success: A boolean indicating whether the response represents a success (true) or failure (false).
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetSuccess(success bool) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Success = success

		return nil
	})

}

// SetMessage assigns a message to the HTTP response options to provide additional context or details.
//
// Parameters:
//   - message: A string representing the message (e.g., success confirmation or error description).
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetMessage(message string) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Message = message

		return nil
	})

}

// SetCode assigns a status or custom code to the HTTP response.
//
// Parameters:
//   - code: The code to set for the response, as defined by type parameter C (can be int or string).
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetCode(code C) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Code = code

		return nil
	})

}

// SetData assigns the main content or payload to the HTTP response options.
//
// Parameters:
//   - data: The data to include in the response, defined by type parameter D, which can be any type.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetData(data D) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Data = data

		return nil
	})

}

// SetExtra adds additional metadata to the HTTP response options.
//
// Parameters:
//   - extra: A map of additional metadata defined by type E, providing supplementary details beyond standard fields.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetExtra(extra E) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Extra = extra

		return nil
	})

}

// SetTotal assigns a total count or amount to the HTTP response, often used for pagination or summaries.
//
// Parameters:
//   - total: A value representing the total, as defined by integer type parameter T.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetTotal(total T) {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Total = total

		return nil
	})

}

// List provides the list of option functions used to configure the HTTP response.
//
// Returns:
//   - []func(*HTTPResponseOptions[C, D, E, T]) error: A slice of functions that configure the response options.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) List() []func(*HTTPResponseOptions[C, D, E, T]) error {
	return httpResponseBuilder.Opts
}
