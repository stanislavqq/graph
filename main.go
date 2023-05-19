package main

import (
	"graph/app"
	"graph/server"
)

func main() {
	app.NewChart()
	server.New().Start()
}
