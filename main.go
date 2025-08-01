package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [server|client]")
		return
	}

	switch os.Args[1] {
	case "server":
		RunServer()
	case "client":
		RunClient()
	default:
		fmt.Println("Invalid option. Use 'server' or 'client'")
	}
}
