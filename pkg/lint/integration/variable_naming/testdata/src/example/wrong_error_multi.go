package example

import "fmt"

func WrongErrorMulti() {
	fail := fmt.Errorf("test")
	_ = fail
}
