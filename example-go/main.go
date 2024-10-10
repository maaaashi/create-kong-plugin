package main

import (
	"github.com/Kong/go-pdk/server"
	"github.com/yourusername/my-go-plugin/handler"
)

func main() {
	server.StartServer(handler.New, Version, Priority)
}

const Version = "1.0.0"
const Priority = 10
