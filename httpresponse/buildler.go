// Package httpresponse offers a builder-pattern implementation to construct flexible, reusable HTTP responses.
// This builder simplifies setting key response attributes, including status, message, code, data, metadata, and total count,
// for consistent and customizable HTTP responses across applications.
package httpresponse

// HTTPResponseBuilder is a generic builder for constructing structured HTTP response configurations.
// It allows setting various response fields such as success status, message, response code, data, total count, and additional metadata.
//
// Type parameters:
//   - C: Defines the type for the response code, supporting either int or string.
//   - D: Defines the type for the data field, which can be any data type (e.g., string, struct, array, etc.).
//   - E: Defines the type for extra metadata, represented as a map with string keys and any values.
//   - T: Defines the type for the total field, supporting various integer types (e.g., int, uint, int64).
type HTTPResponseBuilder[
	C int | string,
	D any,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
] struct {
	Opts []func(*HTTPResponseOptions[C, D, E, T]) error
}

// HTTPResponse initializes a new instance of HTTPResponseBuilder with default settings.
// By default, the Success field is set to true, indicating a successful response.
//
// Returns:
//   - *HTTPResponseBuilder: An instance of HTTPResponseBuilder with default success status.
func HTTPResponse[
	C int | string,
	D any,
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

// SetSuccess specifies the Success field in the HTTP response options.
//
// Parameters:
//   - success: A boolean indicating whether the response is successful (true) or not (false).
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetSuccess(success bool) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Success = success

		return nil
	})

	return httpResponseBuilder
}

// SetMessage adds a message to the HTTP response options for providing additional context or detail.
//
// Parameters:
//   - message: A string containing the message, such as a success confirmation or error description.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetMessage(message string) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Message = message

		return nil
	})

	return httpResponseBuilder
}

// SetData includes the main content or payload in the HTTP response options.
//
// Parameters:
//   - data: The content to include in the response, defined by type parameter D, which can be any type.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetCode(code C) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Code = code

		return nil
	})

	return httpResponseBuilder
}

// SetData assigns the main content or payload to the HTTP response options.
//
// Parameters:
//   - data: The data to include in the response, defined by type parameter D, which can be any type.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetData(data D) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Data = data

		return nil
	})

	return httpResponseBuilder
}

// SetExtra adds supplementary metadata to the HTTP response options.
//
// Parameters:
//   - extra: A map of additional metadata, defined by type parameter E, for providing extra details beyond standard fields.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetExtra(extra E) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Extra = extra

		return nil
	})

	return httpResponseBuilder
}

// SetTotal specifies a total count or amount in the HTTP response, typically used for pagination or summaries.
//
// Parameters:
//   - total: The total value, defined by integer type parameter T.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) SetTotal(total T) *HTTPResponseBuilder[C, D, E, T] {
	httpResponseBuilder.Opts = append(httpResponseBuilder.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Total = total

		return nil
	})

	return httpResponseBuilder
}

// List retrieves the list of option functions that configure the HTTP response.
//
// Returns:
//   - []func(*HTTPResponseOptions[C, D, E, T]) error: A slice of functions used to configure the response options.
func (httpResponseBuilder *HTTPResponseBuilder[C, D, E, T]) List() []func(*HTTPResponseOptions[C, D, E, T]) error {
	return httpResponseBuilder.Opts
}
