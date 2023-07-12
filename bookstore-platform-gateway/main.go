package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	RegisterRoute(h)

	h.Spin()
}
