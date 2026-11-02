package deletion

type EmptyCheck struct {
	Count   func(string) (int64, error)
	Message string
}
