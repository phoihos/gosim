package database

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	// Importing mssql driver package only for gorm
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// A SetupOption is used to generate database Configuration
type SetupOption struct {
	Alias    string
	Host     string
	Port     string
	Database string // dbName
	User     string
	Password string
}

// A Configuration is used to open database
type Configuration interface {
	GetAlias() string
	BuildConnString() (string, string) // engineName, connectionString
}

var pool map[string]*gorm.DB

// Close all database connections
func Close() {
	for _, db := range pool {
		db.Close()
	}
}

// OpenConnection open a database connection
func OpenConnection(config Configuration) error {
	dbConnection, err := gorm.Open(config.BuildConnString())
	if err != nil {
		return err
	}

	dbConnection.DB().SetMaxIdleConns(10)
	dbConnection.DB().SetMaxOpenConns(100)
	dbConnection.DB().SetConnMaxLifetime(time.Hour)

	pool[strings.ToLower(config.GetAlias())] = dbConnection

	return nil
}

// GetConnection return a database connection by alias
func GetConnection(alias string) *gorm.DB {
	return pool[strings.ToLower(alias)]
}

func init() {
	// Change "Column Naming Rule"
	gorm.TheNamingStrategy.Column = func(name string) string { return name }

	pool = make(map[string]*gorm.DB)
}
