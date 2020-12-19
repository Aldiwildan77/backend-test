package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PgDatabase to store database connection
type PgDatabase struct {
	db *gorm.DB
}

// NewInstance to create a pool connection
func (pg *PgDatabase) NewInstance(dsn string) (*gorm.DB, error) {
	instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return instance, nil
}
