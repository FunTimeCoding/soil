package user

import (
	"other.test/lib"
	"other.test/lib/maker"
)

func Use() string {
	return lib.Log("a")
}

func Width() int {
	return maker.Make().Width
}
