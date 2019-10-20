package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type OnePassword struct {
	path string
}

type ImportStruct struct {
	Title    string
	Website  string
	Username string
	Password string
	Notes    string
}

func NewOnePassword(path *string) *OnePassword {
	return &OnePassword{
		path: *path,
	}
}

func (onePassword *OnePassword) Convert(enpass *Export) *[][]string {
	result := make([][]string, 0, len(enpass.Items))

	result = append(result, []string{"title", "website", "username", "password", "notes"})

	for _, item := range enpass.Items {
		var record []string

		var website string
		var username string
		var password string
		var notes string

		fields := make(map[string][]string, 0)
		for _, field := range item.Fields {
			if field.Value == "" {
				continue
			}

			if fields[field.Type] == nil {
				row := []string{field.Value}
				fields[field.Type] = row
			} else {
				fields[field.Type] = append(fields[field.Type], field.Value)
			}
		}

		website = appendValue(fields["url"])
		password = appendValue(fields["password"])

		if len(fields["email"]) > 0 {
			username = appendValue(fields["email"])

			if len(appendValue(fields["username"])) > 0 {
				notes = fmt.Sprintf("username: %s", appendValue(fields["username"]))
			}
		} else if len(appendValue(fields["username"])) > 0 {
			username = appendValue(fields["username"])
		}

		if item.Note != "" {
			notes = notes + "\n" + item.Note
		}

		record = append(record, item.Title, website, username, password, notes)

		result = append(result, record)
	}

	return &result
}

func (onePassword *OnePassword) ToCsv(records *[][]string) {
	file, err := os.Create(onePassword.path)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.WriteAll(*records)
	if err != nil {
		panic(err)
	}
}

func appendValue(source []string) (result string) {
	for _, v := range source {
		if result == "" {
			result = v
		} else {
			result = result + ", " + v
		}
	}

	return result
}
