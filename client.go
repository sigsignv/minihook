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

func (c *Client) LatestEntryID() (int64, error) {
	f := &miniflux.Filter{
		Order:     "id",
		Direction: "desc",
		Limit:     1,
	}

	r, err := c.queryEntries(f)
	if err != nil {
		return -1, err
	}

	for _, e := range r {
		return e.ID, nil
	}

	return -1, fmt.Errorf("miniflux has no entry")
}

func (c *Client) NewEntries(entryID int64) ([]Entry, error) {
	f := &miniflux.Filter{
		Order:        "id",
		Status:       "unread",
		AfterEntryID: entryID,
	}

	r, err := c.queryEntries(f)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for _, e := range r {
		entry := Entry{
			ID:      e.ID,
			Title:   e.Title,
			URL:     e.URL,
			Date:    e.Date,
			Content: e.Content,
			Author:  e.Author,
		}
		entries = append(entries, entry)
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

func (e Entry) String() string {
	return fmt.Sprintf("%d: \"%s\" (%s)", e.ID, e.Title, e.URL)
}
