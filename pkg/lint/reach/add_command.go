package reach

func addCommand(
	result *Edits,
	files map[string]bool,
	w *CommandWrites,
) {
	for _, p := range w.Paths {
		files[p] = true
	}

	switch {
	case w.Fixer:
		result.Fixers++
	case w.Script:
		result.Scripts++
	case w.Unresolved:
		result.Unresolved++
	case len(w.Paths) > 0:
		result.Resolved++
	}
}
