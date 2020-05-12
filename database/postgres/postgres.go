package mysql

import (
	"fmt"

	"github.com/phoihos/gosim/database"

	// Importing driver package only for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func buildConnString(config *database.Configuration) (alias string, connString string) {
	return "postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

func init() {
	database.BuildConnString = buildConnString
}
