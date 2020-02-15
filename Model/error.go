package Model

type ErrorMessage struct {
	Type    string
	Message string
}

type ErrorResponse struct {
	Error []ErrorMessage
}
