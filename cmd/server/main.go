package main

import (
	"github.com/dev3mike/go-api-swagger-boilerplate/cmd/server/setup"
	_ "github.com/dev3mike/go-api-swagger-boilerplate/docs"
)

// @title 4d Backend API
// @version 1.0
// @description This is the API for the 4d project backend
func main() {
	setup.SetupServerPrerequisites()
	setup.StartServer()
}
