package main

import "go-driver-register/internal/server"

// @title Go Driver Register API
// @version 1.0
// @description API for managing drivers and vehicles
// @host localhost:8000
// @basePath /
func main() {
	server.StartServer()
}
