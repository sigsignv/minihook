package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Position struct {
	ID int64 `json:"id"`
}

func NewPosition(r io.Reader) (*Position, error) {
	var p Position

	d := json.NewDecoder(r)
	if err := d.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func LoadPosition(path string) (*Position, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Position{ID: -1}, nil
		}
		return nil, err
	}
	defer f.Close()

	return NewPosition(f)
}

func (p *Position) IsIncreased(latest *Position) bool {
	return p.ID < latest.ID
}

func (p *Position) Save(w io.Writer) error {
	e := json.NewEncoder(w)
	if err := e.Encode(p); err != nil {
		return err
	}

	return nil
}

func (p *Position) SaveFile(path string) error {
	dir, file := filepath.Split(path)
	f, err := os.CreateTemp(dir, file)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := p.Save(f); err != nil {
		return err
	}

	if err := os.Rename(f.Name(), path); err != nil {
		return err
	}

	return nil
}
