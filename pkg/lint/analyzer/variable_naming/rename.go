package variable_naming

import "go/types"

type Rename struct {
	Object  types.Object
	NewName string
}
