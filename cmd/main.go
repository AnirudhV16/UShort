package main

import (
	"UShort/cmd/api"
	"UShort/types"
	"log"
)

func main() {
	store := make(types.Mystore)
	//counter value is stored inside store itself as a global value to increment
	store["counter"] = "1"

	server := api.NewAPIServer(":8080", store)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
