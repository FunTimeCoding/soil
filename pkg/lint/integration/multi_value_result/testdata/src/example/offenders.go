package example

import "fmt"

func pair() (string, int) {
	return "value", 200
}

func Offenders() {
	pair()
	fmt.Println(pair())
	go pair()

	defer pair()
}
