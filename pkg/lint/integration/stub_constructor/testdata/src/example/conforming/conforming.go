package conforming

type Record struct {
	Name string `gorm:"index"`
}

type Filter struct {
	Limit int
}

func Stub() *Record {
	return &Record{}
}

func New(name string) *Record {
	return &Record{Name: name}
}

func NewFilter() *Filter {
	return &Filter{}
}
