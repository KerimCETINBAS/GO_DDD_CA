package auth_contract

type AuthRegisterResponse struct {
	Id       string `json:"_id"`
	UserName string `json:"userName"`
}

type AuthLoginResponse struct {
	Id      string `json:"_id"`
	TraceId string `json:"_traceId"`
}
