package mock_system

func New() *System {
	return &System{
		outputs:  make(map[string]string),
		failures: make(map[string]error),
	}
}
