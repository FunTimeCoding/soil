package example

import "fmt"

func WrongErrorSingleLetter() {
	x := fmt.Errorf("test")
	_ = x
}
