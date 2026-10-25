package example

func Clean() []*Service {
	return []*Service{
		NewService(reporter{}, client{}, nil, reporter{}),
		NewService(reporter{}, client{}, nil, nil...),
		NewHolder(nil),
		Build(nil),
	}
}
