package models

import (
	"time"
)

//Base ...
type Base struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	DeletedBy string    `json:"deletedBy,omitempty"`
}
