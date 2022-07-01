package models

import "time"

type Notifications struct {
	ID        int64 `gorm:"primary_key"`
	UserID    int64 ``
	ProjectID int64 ``

	ReadAt *time.Time `json:"read_at" gorm:"column:read_at" sql:"not null;type:datetime"`
	SendAt *time.Time `json:"send_at" gorm:"column:send_at" sql:"not null;type:datetime"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"not null;type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"not null;type:datetime"`
	DeletedAt *time.Time
}
