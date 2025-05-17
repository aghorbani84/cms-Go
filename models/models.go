package models



// Response is a generic response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}



// NewResponse creates a new response with the given parameters
func NewResponse(success bool, message string, data interface{}) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}