package service

type countTarget struct {
	read   func(string) (int64, error)
	target *int64
}
