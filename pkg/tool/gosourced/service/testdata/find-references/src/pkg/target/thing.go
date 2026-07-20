package target

type Thing struct{}

func (t *Thing) Ping() string {
	return "charlie"
}
