package main

import (
	"fmt"

	miniflux "miniflux.app/client"
)

type Client struct {
	Server string
	Token  string
}

func NewClient(config *Config) (*Client, error) {
	if config.Server == "" {
		return nil, fmt.Errorf("[Client] Server is empty")
	}
	if config.Token == "" {
		return nil, fmt.Errorf("[Client] Token is empty")
	}

	c := &Client{
		Server: config.Server,
		Token:  config.Token,
	}

	return c, nil
}

func (c *Client) LatestEntryID() (*Position, error) {
	f := &miniflux.Filter{
		Order:     "id",
		Direction: "desc",
		Limit:     1,
	}

	r, err := c.queryEntries(f)
	if err != nil {
		return nil, err
	}

	for _, e := range r {
		return &Position{ID: e.ID}, nil
	}

	return nil, fmt.Errorf("miniflux has no entry")
}

func (c *Client) NewEntries(p *Position) ([]Entry, error) {
	f := &miniflux.Filter{
		Order:        "id",
		Status:       "unread",
		AfterEntryID: p.ID,
	}

	r, err := c.queryEntries(f)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for _, e := range r {
		entry, err := NewEntry(e)
		if err != nil {
			continue
		}
		entries = append(entries, *entry)
	}

	return entries, nil
}

func (c *Client) queryEntries(filter *miniflux.Filter) (miniflux.Entries, error) {
	m := miniflux.New(c.Server, c.Token)

	r, err := m.Entries(filter)
	if err != nil {
		return nil, err
	}

	return r.Entries, nil
}
