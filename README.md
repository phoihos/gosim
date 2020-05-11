# Go Simple Webserver

This Go package is a simple webserver framework.

## Importing

`import github.com/phoihos/gosim`

## Usage

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
	"io"
	"net/http"

	"github.com/phoihos/gosim/route"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func init() {
	route.MapRouteFunc("/", handle)
}
```

## License

Released under the [MIT License](LICENSE)
