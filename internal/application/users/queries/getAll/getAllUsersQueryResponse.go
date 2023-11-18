package users_getall_query

type GetUserQueryResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type GetAllUsersQueryResponse []GetUserQueryResponse
