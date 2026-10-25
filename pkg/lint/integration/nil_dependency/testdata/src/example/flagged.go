package example

func Flagged() []*Service {
	return []*Service{
		NewService(nil, nil, nil),
		NewService(reporter{}, client{}, nil, reporter{}, nil),
	}
}
