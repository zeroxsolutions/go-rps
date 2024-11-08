package httpresponse

type HTTPResponseBuilder[
	C int | string,
	D any,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
] struct {
	Opts []func(*HTTPResponseOptions[C, D, E, T]) error
}

func HTTPResponse[
	D any,
	C int | string,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
	E map[string]any,
]() *HTTPResponseBuilder[C, D, E, T] {

	httpResponse := new(HTTPResponseBuilder[C, D, E, T])

	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Success = true

		return nil
	})

	return httpResponse
}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetSuccess(success bool) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Success = success

		return nil
	})

}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetMessage(message string) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Message = message

		return nil
	})

}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetCode(code C) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Code = code

		return nil
	})

}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetData(data D) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Data = data

		return nil
	})

}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetExtraResponse(extraResponse E) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.ExtraResponse = extraResponse

		return nil
	})

}

func (httpResponse *HTTPResponseBuilder[C, D, E, T]) SetTotal(total T) {
	httpResponse.Opts = append(httpResponse.Opts, func(args *HTTPResponseOptions[C, D, E, T]) error {

		args.Total = total

		return nil
	})

}
