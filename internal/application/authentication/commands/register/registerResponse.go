package register_command

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserCommandResponse struct {
	Id        uuid.UUID `json:"id"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
