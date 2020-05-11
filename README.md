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
	"github.com/phoihos/gosim/database/mssql"
	"github.com/phoihos/gosim/server"

	_ "hello"
)

func main() {
	defer database.Close()

	dbOpt := &database.SetupOption{Alias: "example", Host: "127.0.0.1", Port: "1433", Database: "exam", User: "user", Password: "password"}
	dbConf := mssql.NewConfiguration(dbOpt)
	if err := database.OpenConnection(dbConf); err != nil {
		log.Fatal(err)
	}

	conf := &server.Configuration{Port: "8080", ShutdownPath: "/shutdown"}
	server.Run(conf)
}
```

The hello handler code is :
```go
package hello

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/phoihos/gosim/database"
	"github.com/phoihos/gosim/route"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
    db := database.GetConnection("example")

    type product struct {
        Code string
        Price uint
    }
    var results []product
    db.Raw("select * form products").Scan(&results)
    
    b, _ := json.Marshal(results)
	jsonText := string(b)
	io.WriteString(w, jsonText)
}

func init() {
    route.MapRouteFunc("/", handle)
    route.MapRouteFunc("/products", handleProducts)
}
```

## License

Released under the [MIT License](LICENSE)
