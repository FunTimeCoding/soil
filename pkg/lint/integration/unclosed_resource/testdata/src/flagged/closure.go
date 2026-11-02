package example

func Goroutine() {
	c := newClient()
	go func() { c.Work() }()
}

func LocalClosure() {
	c := newClient()
	work := func() { c.Work() }
	work()
}

func register(func()) {}

func UnrelatedClosure() {
	c := newClient()
	register(func() {})
	c.Work()
}
