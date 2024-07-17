package main

import "github.com/kikemaru/duiroPlatform/internal/app"

// @title Duirole platform API
// @version 1.0
// @description API Server for platform telegram bot [duirole]

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
