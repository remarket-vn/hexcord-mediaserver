package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/remarket-vn/hexcord-mediaserver/lib"
)

func main() {
	lib.InitializeRoutes()

	fmt.Printf("Starting server at port 8090\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
