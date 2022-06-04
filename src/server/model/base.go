package model

import "time"

type Base struct {
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt NullTime
	UpdatedBy string
}
