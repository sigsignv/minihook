package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Position struct {
	ID int64 `json:"id"`
}

func LoadPosition(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return -1, nil
		}
		return -1, err
	}
	defer f.Close()

	var p Position
	d := json.NewDecoder(f)
	if err := d.Decode(&p); err != nil {
		return -1, err
	}

	return p.ID, nil
}

func SavePosition(path string, id int64) error {
	dir, file := filepath.Split(path)
	f, err := os.CreateTemp(dir, file)
	if err != nil {
		return err
	}
	defer f.Close()

	p := Position{
		ID: id,
	}

	e := json.NewEncoder(f)
	if err := e.Encode(p); err != nil {
		return err
	}

	if err := os.Rename(f.Name(), path); err != nil {
		return err
	}

	return nil
}
