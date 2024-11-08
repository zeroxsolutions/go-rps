package httpresponse

type HTTPResponseOptions[
	C int | string,
	D any,
	E map[string]any,
	T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
] struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Code          C      `json:"code,omitempty"`
	Data          D      `json:"data,omitempty"`
	Total         T      `json:"total,omitempty"`
	ExtraResponse E      `json:"-"`
}

func (httpResponseOptions *HTTPResponseOptions[C, D, E, T]) MarshalJSON() ([]byte, error) {
	return nil, nil
}
