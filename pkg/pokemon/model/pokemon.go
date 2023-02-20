package model

import (
	"database/sql"
	"time"
)

type Pokemon struct {
	ID        uint         `gorm:"primarykey,autoIncrement" json:"id"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at,omitempty"`

	Name string `gorm:"not null;unique" json:"name"`
}
