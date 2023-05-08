package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func newClientFromFile(filepath string) (*Client, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c, err := LoadConfig(f)
	if err != nil {
		return nil, err
	}

	return NewClient(c)
}

func main() {
	cli := flag.NewFlagSet("minihook", flag.ExitOnError)
	var (
		c = cli.String("c", "./minihook.toml", "Set config location")
		p = cli.String("p", "./position.json", "Set position file location")
		v = cli.Bool("v", false, "Show version")
	)
	cli.Parse(os.Args[1:])

	if *v {
		fmt.Println("v0.1.0")
		os.Exit(0)
	}

	client, err := newClientFromFile(*c)
	if err != nil {
		log.Fatal(err)
	}

	prev, err := LoadPosition(*p)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := client.LatestEntryID()
	if err != nil {
		log.Fatal(err)
	}

	if prev != -1 && cur > prev {
		r, err := client.NewEntries(prev)
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range r {
			fmt.Printf("%v\n", e)
		}
	}

	err = SavePosition(*p, cur)
	if err != nil {
		log.Fatal(err)
	}
}
