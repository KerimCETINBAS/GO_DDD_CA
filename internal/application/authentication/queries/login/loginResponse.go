package login_query

import (
	"github.com/google/uuid"
)

type LoginUserQueryResponse struct {
	Id           uuid.UUID
	UserName     string
	Email        string
	AccessToken  string
	RefreshToken string
}
