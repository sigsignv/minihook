package main

import (
	"fmt"
	"time"

	miniflux "miniflux.app/client"
)

type Client struct {
	Server string
	Token  string
}

type Entry struct {
	ID      int64
	Title   string
	URL     string
	Date    time.Time
	Content string
	Author  string
}

func NewClient(c *Config) (*Client, error) {
	if c.Server == "" {
		return nil, fmt.Errorf("server is empty")
	}
	if c.Token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	client := &Client{
		Server: c.Server,
		Token:  c.Token,
	}
	return client, nil
}

func (c *Client) LatestEntryID() (int64, error) {
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

func (c *Client) NewEntries(entryID int64) ([]Entry, error) {
	filter := &miniflux.Filter{
		Order:        "id",
		Status:       "unread",
		AfterEntryID: entryID,
	}
	entries, err := c.queryEntries(filter)
	if err != nil {
		return nil, err
	}
	var results []Entry
	for _, entry := range entries {
		r := Entry{
			ID:      entry.ID,
			Title:   entry.Title,
			URL:     entry.URL,
			Date:    entry.Date,
			Content: entry.Content,
			Author:  entry.Author,
		}
		results = append(results, r)
	}
	return results, nil
}

func (c *Client) Verify() error {
	client := miniflux.New(c.Server, c.Token)
	_, err := client.Me()
	return err
}

func (c *Client) queryEntries(filter *miniflux.Filter) (miniflux.Entries, error) {
	client := miniflux.New(c.Server, c.Token)
	r, err := client.Entries(filter)
	if err != nil {
		return nil, err
	}
	return r.Entries, nil
}
