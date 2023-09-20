package utils

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

func SuccessResponse(message string, data ...interface{}) TSuccessResponse {
	var responseData interface{}

	if len(data) == 0 {
		responseData = nil
	}

	return TSuccessResponse{
		Meta: TResponseMeta{
			Success: true,
			Message: message,
		},
		Results: responseData,
	}
}

func ErrorResponse(message string) TErrorResponse {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}
