package example

func FunctionLocalType() any {
	type entry struct {
		name  string
		count int
	}

	return entry{}
}
