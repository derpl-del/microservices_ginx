package pgadapter

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,         // Disable color
	},
)

type Adapter struct {
	Table      *gorm.DB
	Connection *sql.DB
}

func (adp Adapter) New() Adapter {
	db_url := os.Getenv("DATABASE_URL")
	db_name := os.Getenv("DATABASE_NAME")
	if len(db_url) == 0 {
		db_url = "postgres://postgres:welcome1@localhost:5432/"
		db_name = "banana"
	}
	db, err := gorm.Open(postgres.Open(db_url+db_name), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	connection, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return Adapter{Table: db, Connection: connection}
}
