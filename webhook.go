package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Webhook struct {
	URL string
}

func (w *Webhook) Post(e *Entry) error {
	t, err := formatRFC3339(e.Date)
	if err != nil {
		return err
	}

	body := map[string]string{
		"id":      fmt.Sprintf("%d", e.ID),
		"title":   e.Title,
		"url":     e.URL,
		"date":    t,
		"content": e.Content,
		"author":  e.Author,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(w.URL, "application/json", bytes.NewReader(j))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func formatRFC3339(t time.Time) (string, error) {
	b, err := t.MarshalText()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
