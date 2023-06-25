package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Webhook struct {
	URL string `toml:"url"`
}

func (w *Webhook) Post(e *Entry) error {
	j, err := json.Marshal(e)
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
