package main

import (
	"geia.us/service/server"
)

func main() {
	server.Start(":50051", &server.UserManagerServer{})
}
