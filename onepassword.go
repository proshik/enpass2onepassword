package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	OnePasswordTitle    = "title"
	OnePasswordWebsite  = "website"
	OnePasswordUsername = "username"
	OnePasswordPassword = "password"
	OnePasswordNotes    = "notes"
)

type OnePassword struct {
	path string
}

type Import struct {
	logins [][]string
}

func NewOnePassword(path *string) *OnePassword {
	return &OnePassword{
		path: *path,
	}
}

func (onePassword *OnePassword) Convert(enpass *Export) *Import {
	logins := make([][]string, 0)

	logins = append(logins, []string{OnePasswordTitle, OnePasswordWebsite, OnePasswordUsername, OnePasswordPassword, OnePasswordNotes})

	for _, item := range enpass.Items {
		var website string
		var username string
		var password string
		var notes string

		// build the map type -> slice of values
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

		website = appendValue(fields[EnpassTypeWebsite])
		password = appendValue(fields[EnpassTypePassword])

		if len(fields[EnpassTypeEmail]) > 0 {
			username = appendValue(fields[EnpassTypeEmail])

			if len(appendValue(fields[EnpassTypeUsername])) > 0 {
				notes = fmt.Sprintf("username: %s", appendValue(fields[EnpassTypeUsername]))
			}
		} else if len(appendValue(fields[EnpassTypeUsername])) > 0 {
			username = appendValue(fields[EnpassTypeUsername])
		}

		if item.Note != "" {
			notes = notes + "\n" + item.Note
		}

		logins = append(logins, []string{item.Title, website, username, password, notes})
	}

	return &Import{logins}
}

func (onePassword *OnePassword) ToCsv(importStruct *Import) {
	file, err := os.Create(onePassword.path)
	if err != nil {
		log.Fatalf("error during creation of file to import by path: %s", onePassword.path)
	}

	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.WriteAll(importStruct.logins)
	if err != nil {
		log.Fatalf("error while writing data to *.csv file by path: %s", onePassword.path)
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
