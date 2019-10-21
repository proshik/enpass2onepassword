package main

import (
	"encoding/json"
	"io/ioutil"
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

func (e *Enpass) fromJson() *Export {
	file, err := ioutil.ReadFile(e.path)
	if err != nil {
		panic(err)
	}

	var enpass = Export{}
	err = json.Unmarshal(file, &enpass)
	if err != nil {
		panic(err)
	}

	return &enpass
}

func (e *Enpass) fromCsv() *Export {
	panic("not implemented")
}
