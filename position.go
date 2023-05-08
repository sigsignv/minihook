package main

import (
	"encoding/json"
	"os"
)

type Position struct {
	ID int64 `json:"id"`
}

func LoadPosition(filepath string) (int64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return -1, nil
		}
		return -1, err
	}
	defer f.Close()

	var p Position
	d := json.NewDecoder(f)
	err = d.Decode(&p)
	if err != nil {
		return -1, err
	}

	return p.ID, nil
}

func SavePosition(filepath string, id int64) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	p := Position{
		ID: id,
	}

	e := json.NewEncoder(f)
	err = e.Encode(p)
	if err != nil {
		return err
	}
	return nil
}
