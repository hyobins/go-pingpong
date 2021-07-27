package main

import (
	"github.com/joho/godotenv"
	"os"
)

// Read Config from go_profile file
func init() {
	godotenv.Load(".env.profile")
	godotenv.Load(".env."+os.Getenv("GO_PROFILE"))
}


func main() {
	//db 생성
	db := database.SqlStore{}.GetDb()
	defer db.Close()

	server.Server{MainDB: db}.Init()
}