package source

type Widget struct {
	Size int
}

func (w *Widget) Run() int {
	return w.Size
}
