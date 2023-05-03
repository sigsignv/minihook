package main

import (
	"fmt"
	"log"
	"os"

	miniflux "miniflux.app/client"
)

func main() {
	endpoint, ok := os.LookupEnv("MINIFLUX_ENDPOINT")
	if !ok {
		log.Fatalln("Should set 'MINIFLUX_ENDPOINT'")
	}
	token, ok := os.LookupEnv("MINIFLUX_TOKEN")
	if !ok {
		log.Fatalln("Should set 'MINIFLUX_TOKEN'")
	}
	client := miniflux.New(endpoint, token)
	me, err := client.Me()
	if err != nil {
		log.Fatalln("Error: Login failed.")
	}
	fmt.Println(me)
}
