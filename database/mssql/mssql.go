package mssql

import (
	"fmt"

	"github.com/phoihos/gosim/database"
)

func buildConnString(config *database.Configuration) (alias string, connString string) {
	return "mssql", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

func init() {
	database.BuildConnString = buildConnString
}
