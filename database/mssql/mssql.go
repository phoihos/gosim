package mssql

import (
	"fmt"

	"github.com/phoihos/gosim/database"
)

type config struct {
	option *database.SetupOption
}

func (c *config) GetAlias() string {
	return c.option.Alias
}

func (c *config) BuildConnString() (string, string) {
	return "mssql", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		c.option.User, c.option.Password, c.option.Host, c.option.Port, c.option.Database)
}

// NewConfiguration make configuration for mssql
func NewConfiguration(option *database.SetupOption) database.Configuration {
	return &config{option: option}
}
