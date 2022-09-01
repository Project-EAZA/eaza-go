package common

type Response[T any] struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Data    T      `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Err     error  `json:"err"`
}

func NewSuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Code:    200,
		Msg:     "Success",
		Data:    data,
		Success: true,
	}
}

func NewErrorResponse(code int, err error) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Msg:     "Error",
		Err:     err,
		Success: false,
	}
}
