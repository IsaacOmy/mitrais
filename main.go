package main

import (
	"fmt"

	"github.com/IsaacOmy/mitrais/db"
	"github.com/IsaacOmy/mitrais/handle"
)

func main() {
	db.ConfigureDB()
	fmt.Println("Running on port 8080")
	handle.HandleRequest()
}
