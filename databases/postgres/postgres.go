package postgres

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	psqlMaxRetry = 60
)

// Connector - PSQL connector interface API.
type Connector interface {
	ConnectionString() string
}

// Open connection to postgres server.
func Open(conn Connector) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", conn.ConnectionString())
	if err != nil {
		for i := 1; i <= psqlMaxRetry; i++ {
			db, err = gorm.Open("postgres", conn.ConnectionString())
			if err == nil {
				break
			}
			time.Sleep(time.Duration(i) * time.Second)
		}
	}
	if err != nil {
		return nil, err
	}
	db.SingularTable(false)
	// db.DB().SetConnMaxLifetime(time.Duration(config.Configuration.Postgres.ConnMaxLifetime) * time.Hour)
	// db.DB().SetMaxOpenConns(config.Configuration.Postgres.MaxOpenConnections)
	// db.DB().SetMaxIdleConns(config.Configuration.Postgres.MaxIdleConnections)
	return db, nil
}
