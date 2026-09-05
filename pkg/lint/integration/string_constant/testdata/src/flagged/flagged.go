package flagged

import "fmt"

func Register() {
	fmt.Println("name")  // want `string literal "name" has constants constant.Name or constant.NameColumn`
	fmt.Println("query") // want `string literal "query" has constant constant.Query`
}

func MapKey() {
	m := map[string]any{}
	_ = m["name"]  // want `string literal "name" has constants constant.Name or constant.NameColumn`
	m["query"] = 1 // want `string literal "query" has constant constant.Query`
}
