package database

import (
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/utils/load"
)

var db *gorm.DB

type DBout struct {
	DB *gorm.DB
}

// SetDB sets the database connection
func SetDB(database *gorm.DB) {
	db = database
}

type Logger struct {
	Log *zap.Logger
}

func GormLogger(l *zap.Logger) {
}

type ConfigWrapper struct {
	*load.Config
}

func (cfg ConfigWrapper) GetDB() (out DBout, err error) {

	var b strings.Builder
	b.WriteString("host=")
	b.WriteString(cfg.DB_HOST)
	b.WriteString(" user=")
	b.WriteString(cfg.DB_USERNAME)
	b.WriteString(" password=")
	b.WriteString(cfg.DB_PASSWORD)
	b.WriteString(" dbname=")
	b.WriteString(cfg.DB_NAME)
	b.WriteString(" port=")
	b.WriteString(strconv.Itoa(cfg.DB_PORT))
	b.WriteString(" sslmode=disable TimeZone=UTC")
	dsn := b.String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Set the database connection

	out = DBout{DB: db}

	// Migrate the schema

	return
}
