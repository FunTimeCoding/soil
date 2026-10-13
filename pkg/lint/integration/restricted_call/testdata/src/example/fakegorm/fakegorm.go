package fakegorm

func Open(path string) string {
	return path
}

var Shared = "shared"

type Holder struct {
	Shared string
}

func (h Holder) Open(path string) string {
	return path
}
