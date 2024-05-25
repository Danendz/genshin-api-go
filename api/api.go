package api

type APIResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Success bool `json:"success"`
}

func NewApiResponse(message string, data interface{}, success bool) *APIResponse {
	return &APIResponse{
		Message: message,
		Data: data,
		Success: success,
	}
}