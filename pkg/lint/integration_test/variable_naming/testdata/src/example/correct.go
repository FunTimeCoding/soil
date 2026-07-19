package example

import "fmt"

func Correct() {
	e := fmt.Errorf("test")
	s := "hello"
	i := 42
	_ = e
	_ = s
	_ = i
}
