package main

import _ "fmt"

func main() {
	enpass := NewEnpass("enpass_json.json")
	onePassword := NewOnePassword("1password_result.csv")

	enpassStruct := enpass.read()

	// fmt.Println(enpassStruct)

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("----------%d------------", i)
	// 	fmt.Printf("title: %s\n", enpassStruct.Items[i].Title)

	// 	for _, field := range enpassStruct.Items[i].Fields {
	// 		if field.Value != "" {
	// 			fmt.Printf("%s: %s\n", field.Type, field.Value)
	// 		}
	// 	}

	// 	fmt.Printf("note: %s\n\n", enpassStruct.Items[i].Note)
	// }

	onePassword.write(enpassStruct)
}
