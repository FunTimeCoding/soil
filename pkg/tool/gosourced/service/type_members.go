package service

import "go/types"

func typeMembers(o types.Object) []types.Object {
	result := []types.Object{o}
	named, okay := o.Type().(*types.Named)

	if !okay {
		return result
	}

	for i := range named.NumMethods() {
		result = append(result, named.Method(i))
	}

	if structure, isStruct := named.Underlying().(*types.Struct); isStruct {
		for i := range structure.NumFields() {
			result = append(result, structure.Field(i))
		}
	}

	return result
}
