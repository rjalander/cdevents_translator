package main

import (
	"fmt"
	"github.com/cdevents/translator/pkg/api"
)

func main() {
	fmt.Println("Starting default API server")
	api.StartServer()
}
