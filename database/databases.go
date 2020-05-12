package database

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	// Importing mssql driver package only for gorm
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// Configuration is used to open database
type Configuration struct {
	Alias    string
	Host     string
	Port     string
	Database string // dbName
	User     string
	Password string
}

// NamingStrategy represents naming strategies
type NamingStrategy gorm.NamingStrategy

// BuildConnString is function to build connection string
var BuildConnString func(config *Configuration) (alias string, connString string)

var pool map[string]*gorm.DB

// SetOrmNamingStrategy set DB/Table/Column namer function for orm
func SetOrmNamingStrategy(ns *NamingStrategy) {
	gorm.AddNamingStrategy(&gorm.NamingStrategy{DB: ns.DB, Table: ns.Table, Column: ns.Column})
}

// Close all database connections
func Close() {
	for _, db := range pool {
		db.Close()
	}
}

// OpenConnection open a database connection with configuration
func OpenConnection(config *Configuration) error {
	if BuildConnString == nil {
		return errors.New("Can't found imported database driver")
	}

	dbConnection, err := gorm.Open(BuildConnString(config))
	if err != nil {
		return err
	}

	dbConnection.DB().SetMaxIdleConns(10)
	dbConnection.DB().SetMaxOpenConns(100)
	dbConnection.DB().SetConnMaxLifetime(time.Hour)

	pool[strings.ToLower(config.Alias)] = dbConnection

	return nil
}

// GetConnection return a database connection by alias
func GetConnection(alias string) *gorm.DB {
	return pool[strings.ToLower(alias)]
}

func init() {
	pool = make(map[string]*gorm.DB)
}
