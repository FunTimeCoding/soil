package overlay

func New(directory string) *Overlay {
	return &Overlay{Replace: map[string]string{}, directory: directory}
}
