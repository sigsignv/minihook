package main

import (
	"fmt"
	"time"

	miniflux "miniflux.app/client"
)

type Entry struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func NewEntry(m *miniflux.Entry) (*Entry, error) {
	t, err := timeToString(m.Date)
	if err != nil {
		return nil, err
	}

	e := &Entry{
		ID:      m.ID,
		Title:   m.Title,
		URL:     m.URL,
		Date:    t,
		Content: m.Content,
		Author:  m.Author,
	}

	return e, nil
}

func (e Entry) String() string {
	return fmt.Sprintf("%d: \"%s\" (%s)", e.ID, e.Title, e.URL)
}

func timeToString(t time.Time) (string, error) {
	b, err := t.MarshalText()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
