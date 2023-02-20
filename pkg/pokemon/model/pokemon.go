package model

import (
	"database/sql"
	"time"
)

type Pokemon struct {
	ID        uint         `gorm:"primarykey,autoIncrement"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	DeletedAt sql.NullTime `gorm:"index"`

	Name string `gorm:"not null;unique"`
}
