package main

import (
	"github.com/sunnydaytech/geiaus/service/server"
)

func main() {
	server.Start(":50051", &server.UserManagerServer{})
}
