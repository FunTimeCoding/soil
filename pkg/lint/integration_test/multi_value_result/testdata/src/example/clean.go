package example

import "fmt"

func single() string {
	return "value"
}

func emit(text string, code int) {
	fmt.Println(text, code)
}

func Clean() (string, int) {
	text, code := pair()
	emit(text, code)
	single()
	fmt.Println(single())

	return pair()
}
