package example

type LocalStruct struct {
	Value int
}

func SamePackage() LocalStruct {
	return LocalStruct{Value: 1}
}
