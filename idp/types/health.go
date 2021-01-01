package types

type HealthRequest struct {
	Ping string `json:"ping"`
}

type HealthResponse struct {
	RPCResponse
	Pong string `json:"pong"`
}
