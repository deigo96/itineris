package util

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	return Response{Code: "2000", Message: "success", Data: data}
}
