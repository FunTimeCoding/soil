package user

import "other.test/lib"

func Use() string {
	return lib.Log("a", 1)
}
