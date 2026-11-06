package seeded

type Record struct {
	Name string `gorm:"index"`
}

func New() *Record {
	return &Record{Name: "alpha"}
}
