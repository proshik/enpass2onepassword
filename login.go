package main

import (
	"fmt"
	"strings"
)

const LoginType = "login"

const (
	EmailLabel    = "EMAIL"
	Email2Label   = "E-MAIL"
	UsernameLabel = "USERNAME"
	PasswordLabel = "PASSWORD"
	UrlLabel      = "URL"
)

type Login struct {
	categoryType string
}

func (login *Login) Generate(items []EnpassItem) [][]string {
	records := make([][]string, 0)

	records = append(records, []string{"title", "website", "username", "password", "notes"})

	for _, item := range items {
		// build the map type -> slice of values
		fieldValuesByLabel := make(map[string][]string, 0)
		for _, field := range item.Fields {
			// skip the field that contains an empty value
			if field.Value == "" {
				continue
			}

			// set the uppercase to Label because "Enpass" exported values contain different values
			label := strings.ToUpper(field.Label)

			if fieldValuesByLabel[label] == nil {
				fieldValuesByLabel[label] = []string{field.Value}
			} else {
				fieldValuesByLabel[label] = append(fieldValuesByLabel[label], field.Value)
			}
		}

		var username string
		var notes string

		// fill the username by email if the not null. In other case
		if len(fieldValuesByLabel[Email2Label]) > 0 {
			username = joinValue(fieldValuesByLabel[Email2Label])

			if len(joinValue(fieldValuesByLabel[UsernameLabel])) > 0 {
				notes = fmt.Sprintf("username(s): %s;\n", joinValue(fieldValuesByLabel[UsernameLabel]))
			}
		} else if len(joinValue(fieldValuesByLabel[UsernameLabel])) > 0 {
			username = joinValue(fieldValuesByLabel[UsernameLabel])
		} else if len(joinValue(fieldValuesByLabel[EmailLabel])) > 0 {
			username = joinValue(fieldValuesByLabel[EmailLabel])
		}

		website := joinValue(fieldValuesByLabel[UrlLabel])
		password := joinValue(fieldValuesByLabel[PasswordLabel])

		if item.Note != "" {
			notes = notes + item.Note
		}

		records = append(records, []string{item.Title, website, username, password, notes})
	}

	return records
}

func (login *Login) Type() string {
	return login.categoryType
}

func joinValue(source []string) (result string) {
	for _, v := range source {
		if result == "" {
			result = v
		} else {
			result = result + ", " + v
		}
	}

	return result
}
