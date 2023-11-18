package services

import "time"

type IDateTimeProvider interface {
	DateTime() time.Time
}
