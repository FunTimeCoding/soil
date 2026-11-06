package flagged

type Record struct {
	Name string `gorm:"index"`
}

func New() *Record { return &Record{} }
