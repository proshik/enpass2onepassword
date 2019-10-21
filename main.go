package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// read cli params
	enpassFilePath := flag.String("enpass_src_path", "enpass.json", "path to file with exported data from the Enpass application")
	enpassExtension := flag.String("enpass_extension", "json", "extension of source file is JSON or CSV (JSON by default)")
	onePasswordPath := flag.String("1password_dst_path", "1password.csv", "path to file with result of conversion")
	flag.Parse()

	// validate flags
	checkFilePath(enpassFilePath)
	checkCsvOrJsonExtention(enpassExtension)

	// init services
	enpass := NewEnpass(enpassFilePath, enpassExtension)
	onePassword := NewOnePassword(onePasswordPath)

	// read enpass file from json
	enpassStruct := enpass.fromJson()

	// invoke convert method
	var onePasswordImport *Import
	switch *enpassExtension {
	case "json":
		onePasswordImport = onePassword.Convert(enpassStruct)
	case "csv":
		onePasswordImport = onePassword.Convert(enpassStruct)
	default:
		log.Fatalf("unexpected extension of input file %s", *enpassExtension)
	}

	// save to csv for 1Password
	onePassword.ToCsv(onePasswordImport)
}

func checkCsvOrJsonExtention(extension *string) {
	if *extension != "json" && *extension != "csv" {
		log.Fatalf("extension %s is not inapplicable", *extension)
	}
}

func checkFilePath(path *string) {
	info, err := os.Stat(*path)
	if os.IsNotExist(err) || info.IsDir() {
		log.Fatalf("file doesn't exist by path %s", *path)
	}
}
