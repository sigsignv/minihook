package main

import (
	"fmt"

	miniflux "miniflux.app/client"
)

type MinifluxClient struct {
	Server string
	Token  string
}

func (c *MinifluxClient) LatestEntryID() (int64, error) {
	filter := &miniflux.Filter{
		Order:     "id",
		Direction: "desc",
		Limit:     1,
	}
	entries, err := c.queryEntries(filter)
	if err != nil {
		return -1, err
	}
	for _, entry := range entries {
		return entry.ID, nil
	}
	return -1, fmt.Errorf("miniflux has no entry")
}

func (c *MinifluxClient) NewEntries(entryID int64) ([]string, error) {
	filter := &miniflux.Filter{
		Order:        "id",
		Status:       "unread",
		AfterEntryID: entryID,
	}
	entries, err := c.queryEntries(filter)
	if err != nil {
		return nil, err
	}
	var array []string
	for _, entry := range entries {
		array = append(array, entry.Title)
	}
	return array, nil
}

func (c *MinifluxClient) Verify() error {
	client := miniflux.New(c.Server, c.Token)
	_, err := client.Me()
	return err
}

func (c *MinifluxClient) queryEntries(filter *miniflux.Filter) (miniflux.Entries, error) {
	client := miniflux.New(c.Server, c.Token)
	r, err := client.Entries(filter)
	if err != nil {
		return nil, err
	}
	return r.Entries, nil
}
