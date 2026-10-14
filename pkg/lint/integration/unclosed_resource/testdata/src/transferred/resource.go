package example

type Resource struct{}

func open() *Resource {
	return &Resource{}
}

func (r *Resource) Close() {}
