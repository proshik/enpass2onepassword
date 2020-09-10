package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

// change to nested structure
type EnpassExport struct {
	Items []EnpassItem `json:"items,omitempty"`
}

type EnpassItem struct {
	Title    string `json:"title,omitempty"`
	Category string `json:"category,omitempty"`
	Note     string `json:"note,omitempty"`
	Fields   []struct {
		Label string `json:"label"`
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"fields,omitempty"`
}

// generator of different import type of files to import in 1Password
type Generator interface {
	Generate(items []EnpassItem) [][]string
	Type() string
}

var generators = []Generator{&Login{LoginType}, &Login{"computer"}, &CreditCard{}}

func main() {
	// read cli params
	enpassFilePath := flag.String("enpass_src_path", "enpass.json", "path to enpassExport with exported data from the Enpass application")
	flag.Parse()

	// validate flags
	checkFilePath(enpassFilePath)

	// read enpass enpassExport from json
	enpassExport, err := os.Open(*enpassFilePath)
	if err != nil {
		log.Fatalf("can't read the enpassExport by path: %s", *enpassFilePath)
	}

	enpassStruct, err := readJSON(enpassExport)
	if err != nil {
		log.Fatalf("error during to unmarshal the json from file by path: %s", *enpassFilePath)
	}
	log.Println("The Enpass enpassExport was read.")

	// group available items by category
	categories := groupByCategory(enpassStruct)

	// convert items to
	toImport := convert(categories)

	// save to files
	createImport(toImport)

	log.Println("SUCCESSFUL GENERATION FILES")
}

func checkFilePath(path *string) {
	info, err := os.Stat(*path)
	if os.IsNotExist(err) || info.IsDir() {
		log.Fatalf("file %s doesn't exist", *path)
	}
}

func readJSON(reader io.Reader) (*EnpassExport, error) {
	var data = EnpassExport{}
	//err = json.Unmarshal(file, &data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func groupByCategory(export *EnpassExport) map[string][]EnpassItem {
	categories := make(map[string][]EnpassItem, 0)
	for _, item := range export.Items {
		// group items by category title
		items := categories[item.Category]
		if items == nil {
			categories[item.Category] = []EnpassItem{item}
		} else {
			categories[item.Category] = append(categories[item.Category], item)
		}
	}

	return categories
}

// key it is the type of items, value is the slice of slices with values (in fact a CSV structure)
func convert(categories map[string][]EnpassItem) map[string][][]string {
	generatorsByType := make(map[string]Generator, 0)
	for _, g := range generators {
		generatorsByType[g.Type()] = g
	}
	result := make(map[string][][]string, 0)
	for k, v := range generatorsByType {
		result[k] = v.Generate(categories[k])
	}

	return result
}

func createImport(records map[string][][]string) {
	// save to csv for 1Password
	for k, v := range records {
		filename := "1password_" + k + ".csv"

		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error during creation of file to import by path: %s", filename)
		}

		err = writeCsv(file, v)
		if err != nil {
			log.Fatalf("error while writing data to *.csv file by path: %s", filename)
		}
	}
}

func writeCsv(w io.Writer, records [][]string) error {
	writer := csv.NewWriter(w)

	defer writer.Flush()

	err := writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}
