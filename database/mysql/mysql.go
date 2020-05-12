package mysql

import (
	"fmt"

	"github.com/phoihos/gosim/database"

	// Importing driver package only for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func buildConnString(config *database.Configuration) (alias string, connString string) {
	return "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

func init() {
	database.BuildConnString = buildConnString
}
