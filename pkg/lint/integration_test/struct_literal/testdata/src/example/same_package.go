package example

// Local struct - bare &LocalStruct{} must be allowed (same package, no selector).
type LocalStruct struct {
	Value int
}

func SamePackage() *LocalStruct {
	return &LocalStruct{}
}
