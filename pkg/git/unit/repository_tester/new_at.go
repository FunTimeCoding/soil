package repository_tester

func NewAt(clone string) *Tester {
	return &Tester{Clone: clone}
}
