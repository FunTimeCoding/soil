package service

func (s *Service) Log(
	name string,
	all bool,
) ([]string, int, error) {
	return s.supervisor.LogLines(name, all)
}
