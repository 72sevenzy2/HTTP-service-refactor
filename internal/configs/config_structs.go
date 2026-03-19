package configs

type GreetResponse struct {
	Message string `json:"message"`
	Count int `json:"count"`
}

type GreetRequest struct {
	Name string `json:"name"`
}