package types

import (
	"encoding/json"
)

type Bundle []Data

type Data struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func NewBundleFromJSON(data []byte) (Bundle, error) {
	b := &Bundle{}

	if err := json.Unmarshal(data, b); err != nil {
		return nil, err
	}

	return *b, nil
}
