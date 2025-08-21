package model

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
		Error:  "",
	}
}

func ERROR(message string) Response {
	return Response{
		Status: StatusError,
		Error:  message,
	}
}
