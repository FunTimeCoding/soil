package target

import "strings"

// Widget holds a display name.
type Widget struct {
	Name string
}

func (w *Widget) Trimmed() string {
	return strings.TrimSpace(w.Name)
}

func NewWidget(name string) *Widget {
	return &Widget{Name: name}
}
