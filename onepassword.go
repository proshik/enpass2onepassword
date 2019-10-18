package main

import (
	"encoding/csv"
	"fmt"
	"os"
	_ "log"
)

type OnePassword struct {
	path string
}

type Import struct {
	Title    string
	Website  string
	Username string
	Password string
	Notes    string
}

func NewOnePassword(path string) *OnePassword {
	return &OnePassword{
		path: path,
	}
}

func (op *OnePassword) write(enpass *Export) {
	result := make([][]string, 0)

	result = append(result, []string{"title", "website", "username", "password", "notes"})

	for _, item := range enpass.Items {
		var record []string

		var website string
		var username string
		var password string
		var notes string

		// TODO update
		// fields := make(map[string][]string, 0)

		for _, field := range item.Fields {
			switch field.Type {
			case "url":
				website = appendValue(website, field.Value)
			case "password":
				password = appendValue(password, field.Value)
			case "email":
				username = appendValue(username, field.Value)
			}

			if field.Type == "username" && field.Value != "" {
				// if username == "" {
				// 	log.Print(field.Value)
				// 	username = field.Value
				// } else {
					notes = fmt.Sprintf("username: %s", field.Value)
				// }
			}
		}

		if item.Note != "" {
			notes = notes + "\n" + item.Note
		}

		record = append(record, item.Title, website, username, password, notes)

		result = append(result, record)
	}

	file, err := os.Create(op.path)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.WriteAll(result)
	if err != nil {
		panic(err)
	}
}

func appendValue(source, value string) string {
	if source == "" {
		return value
	} else {
		return source + ", " + value
	}
}
