package parameterised

type Record struct {
	Name string
}

func Stub(name string) *Record {
	return &Record{Name: name}
}
