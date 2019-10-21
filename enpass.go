package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	EnpassTypeEmail    = "email"
	EnpassTypeUsername = "username"
	EnpassTypePassword = "password"
	EnpassTypeWebsite  = "url"
)

type Enpass struct {
	path      string
	extension string
}

type Export struct {
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	Title    string  `json:"title,omitempty"`
	Category string  `json:"category,omitempty"`
	Note     string  `json:"note,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Field struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewEnpass(path *string, expansion *string) *Enpass {
	return &Enpass{
		path:      *path,
		extension: *expansion,
	}
}

func (enpass *Enpass) fromJson() *Export {
	file, err := ioutil.ReadFile(enpass.path)
	if err != nil {
		log.Fatalf("can't read the file by path: %s", enpass.path)
	}

	var data = Export{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error during to unmarshal the json from file by path: %s", enpass.path)
	}

	return &data
}

func (enpass *Enpass) fromCsv() *Export {
	panic("not implemented")
}
