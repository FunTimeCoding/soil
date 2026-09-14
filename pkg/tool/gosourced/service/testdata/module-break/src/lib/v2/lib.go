package lib

type Widget struct {
	Label string
}

func (w *Widget) Name() string {
	return w.Label
}

func Make(label string) *Widget {
	return &Widget{Label: label}
}

func Gone() string {
	return "gone"
}
