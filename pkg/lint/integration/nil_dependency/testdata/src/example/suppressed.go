package example

func Suppressed() *Service {
	return NewService(nil, client{}, nil) // goanalyze:ignore
}
