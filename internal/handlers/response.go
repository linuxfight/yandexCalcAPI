package handlers

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Result string `json:"result"`
}
