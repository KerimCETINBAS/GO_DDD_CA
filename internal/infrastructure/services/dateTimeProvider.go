package services

import (
	"time"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/services"
)

type dateTime struct{}

func UseDateTimeProvider() services.IDateTimeProvider {
	return &dateTime{}
}

func (p *dateTime) DateTime() time.Time {

	return time.Now()
}
