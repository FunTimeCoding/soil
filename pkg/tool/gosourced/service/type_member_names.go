package service

import "go/types"

func typeMemberNames(typeObject types.Object) map[string]bool {
	result := make(map[string]bool)
	named, okay := typeObject.Type().(*types.Named)

	if !okay {
		return result
	}

	for i := range named.NumMethods() {
		result[named.Method(i).Name()] = true
	}

	structure, okay := named.Underlying().(*types.Struct)

	if !okay {
		return result
	}

	for i := range structure.NumFields() {
		result[structure.Field(i).Name()] = true
	}

	return result
}
