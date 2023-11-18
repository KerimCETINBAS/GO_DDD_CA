package auth_contract

type AuthRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
