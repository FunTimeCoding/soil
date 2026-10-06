package tainted

type Service struct{}

func (s *Service) Run() {}

type Record struct {
	Name string
}
