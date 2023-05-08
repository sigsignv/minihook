package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	cli := flag.NewFlagSet("minihook", flag.ExitOnError)
	var (
		c = cli.String("c", "./minihook.toml", "Set config location")
		v = cli.Bool("v", false, "Show version")
	)
	cli.Parse(os.Args[1:])

	if *v {
		fmt.Println("v0.1.0")
		os.Exit(0)
	}

	f, err := os.Open(*c)
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
	entries, _ := client.NewEntries(id - 10)
	fmt.Printf("%v\n", entries)
}
