package main

import (
	"fmt"

	miniflux "miniflux.app/client"
)

type MinifluxClient struct {
	Server string
	Token  string
}

func (c *MinifluxClient) LatestEntryID() (int, error) {
	client := miniflux.New(c.Server, c.Token)
	filter := &miniflux.Filter{
		Order:     "id",
		Direction: "desc",
		Limit:     1,
	}
	r, err := client.Entries(filter)
	if err != nil {
		return -1, err
	}
	entries := r.Entries
	for _, entry := range entries {
		return int(entry.ID), nil
	}
	return -1, fmt.Errorf("miniflux has no entry")
}

func (c *MinifluxClient) Verify() error {
	client := miniflux.New(c.Server, c.Token)
	_, err := client.Me()
	return err
}
