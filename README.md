# Go Simple Webserver

This Go package is a simple webserver framework.
- Depends on [jinzhu/gorm](https://github.com/jinzhu/gorm) package.

## Importing

`import github.com/phoihos/gosim`

## Usage

> If you want to see full example, please move on to https://github.com/phoihos/gosim-example.

The main code is :
```go
package main

import (
	"log"

	"github.com/phoihos/gosim/database"
	_ "github.com/phoihos/gosim/database/postgres"
	//_ "github.com/phoihos/gosim/database/mysql"
	//_ "github.com/phoihos/gosim/database/mssql"
	"github.com/phoihos/gosim/server"

	_ "handler"
)

func main() {
	defer database.Close()

	// Change orm naming rule if you want to change
	// database.SetOrmNamingStrategy(...)

	dbConf := &database.Configuration{Alias: "example", Host: "127.0.0.1", Port: "1433", Database: "exam", User: "user", Password: "password"}
	if err := database.OpenConnection(dbConf); err != nil {
		log.Print(err)
	}

	conf := &server.Configuration{Port: "8080", ShutdownPath: "/shutdown"}
	server.Run(conf)
}
```

The handler code is :
```go
package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/phoihos/gosim/database"
	"github.com/phoihos/gosim/route"
)

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Hello World")
	default:
		http.NotFound(w, r)
	}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection("example")
	if db == nil {
		io.WriteString(w, "No database connection exists")
		return
	}

	type product struct {
		Code  string
		Price uint
	}

	var results []product
	db.Raw("select * form products").Scan(&results)

	b, _ := json.Marshal(results)
	jsonText := string(b)
	io.WriteString(w, jsonText)
}

func init() {
	route.MapRouteFunc("/", handle).
		MapRouteFunc("/products", handleProducts)
}
```

## License

Released under the [MIT License](LICENSE)
