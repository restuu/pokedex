package mysql

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(ctx context.Context, uri string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(uri))
}
