package main

import (
	"os"

	"github.com/hyobins/go-pingpong/infrastructure/server"
	"github.com/joho/godotenv"
)

// Read Config from go_profile file
func init() {
	godotenv.Load(".env.profile")
	godotenv.Load(".env." + os.Getenv("GO_PROFILE"))
}

// @title Go-pingpong Swagger API
// @version 1.0.0
// @host localhost:8080
func main() {
	//서버 시작
	server.Init()
}
