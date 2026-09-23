package spacing

func (s *Spacing) concern(
	key string,
	text string,
	number int,
	line string,
) {
	s.report.AddConcern(key, text, s.path, number, line, true)
}
