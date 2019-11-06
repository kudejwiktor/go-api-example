package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kudejwiktor/go-api-example/app/servid"
)

func main() {
	server := servid.NewApp()
	server.Start()
}
