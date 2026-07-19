package example

import "fmt"

func ErrorChain() {
	e := fmt.Errorf("first")
	x := fmt.Errorf("second")
	_ = e
	_ = x
}
