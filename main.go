package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./minihook.toml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := LoadConfig(f)
	if err != nil {
		log.Fatal(err)
	}
	client, err := NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	id, err := client.LatestEntryID()
	if err != nil {
		log.Fatal(err)
	}
	entries, err := client.NewEntries(id - 10)
	fmt.Println("Entries: %v", entries)
}
