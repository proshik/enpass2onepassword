package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Enpass struct {
	path string
}

type Export struct {
	Items []Item 
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

func NewEnpass(path string) *Enpass {
	if path == "" {
		log.Fatal("need to provide the path of enpass file")
	}

	return &Enpass{
		path: path,
	}
}

func (e *Enpass) read() *Export {
	// file, err := os.Open("enpass_json.json")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	file, err := ioutil.ReadFile(e.path)
	if err != nil {
		panic(err)
	}

	log.Println("file enpass was readed")

	var enpass = Export{}
	// json.NewDecoder(file).Decode(&enpass)

	err = json.Unmarshal(file, &enpass)
	if err != nil {
		panic(err)
	}

	log.Println("enpass file was parsed to structure")

	return &enpass
}
