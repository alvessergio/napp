package main

import (
	"fmt"

	"github.com/alvessergio/pan-integrations/routes"
)

func main() {
	fmt.Println("Starting rest server...")
	routes.HandleRequest()
}
