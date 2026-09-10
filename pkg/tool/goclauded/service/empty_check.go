package service

type emptyCheck struct {
	count   func(string) (int64, error)
	message string
}
