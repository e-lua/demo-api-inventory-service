package response

type Response struct {
	Status int         `json:"status,omitempty"`
	Error  Error       `json:"error"`
	Data   interface{} `json:"data"`
}

type ResponseAuth struct {
	Error Error `json:"error"`
	Data  Auth  `json:"data"`
}
