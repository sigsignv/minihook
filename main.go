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
		n = cli.Bool("n", false, "Dry run")
		p = cli.String("p", "./position.json", "Set position file location")
		v = cli.Bool("v", false, "Show version")
	)
	cli.Parse(os.Args[1:])

	if *v {
		fmt.Println("v0.1.0")
		os.Exit(0)
	}

	config, err := LoadConfig(*c)
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	pos, err := LoadPosition(*p)
	if err != nil {
		log.Fatal(err)
	}

	latest, err := client.LatestEntryID()
	if err != nil {
		log.Fatal(err)
	}

	if pos.IsIncreased(latest) {
		r, err := client.NewEntries(pos)
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range r {
			for _, w := range config.Webhook {
				w.Post(&e)
			}
			fmt.Printf("%v\n", e)
		}
	}

	if !*n {
		err = latest.SaveFile(*p)
		if err != nil {
			log.Fatal(err)
		}
	}
}
