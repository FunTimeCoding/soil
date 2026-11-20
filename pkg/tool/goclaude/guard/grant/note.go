package grant

func (g *Grant) Note(line string) {
	g.notes = append(g.notes, line)
}
