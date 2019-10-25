package main

import "github.com/kudejwiktor/go-api-example/app/servid"

func main() {
	server := servid.NewApp()
	server.Start()
}
