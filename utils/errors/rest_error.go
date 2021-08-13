package errors

import "net/http"

type RestErr struct {
	Messgae string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Messgae: message,
		Status:  http.StatusBadRequest,
		Error:   "bad Request",
	}
}

func NotFoundError(messaeg string) *RestErr {
	return &RestErr{
		Messgae: messaeg,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}
