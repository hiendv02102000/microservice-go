package handler

// BaseResponse struct
type BaseResponse struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}
type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
