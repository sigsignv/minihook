package main

import (
	"fmt"
	"log"
	"os"
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
	client := &MinifluxClient{
		Server: endpoint,
		Token:  token,
	}
	id, err := client.LatestEntryID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest entry ID: ", id)
}
