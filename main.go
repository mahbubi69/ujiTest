package main

import (
	"fmt"
	"ujiTest/base"
)

func main() {

	server := base.NewServer("9080")

	fmt.Println("Server running on port 9080...\n")
	server.Routes()
}
