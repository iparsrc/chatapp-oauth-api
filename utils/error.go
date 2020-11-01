package utils

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "error: bad request",
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "error: not found",
	}
}

func InternalServerErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "error: internal server error",
	}
}
