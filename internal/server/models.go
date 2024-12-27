package server

type solveRequest struct {
	Expression string `json:"expression"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Result float64 `json:"result"`
}
