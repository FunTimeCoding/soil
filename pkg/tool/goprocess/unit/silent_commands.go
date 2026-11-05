package unit

func silentCommands() [][]string {
	return [][]string{
		{"run", "start", "bravo"},
		{"run", "stop", "alfa"},
		{"run", "restart", "alfa"},
		{"run", "reload-procfile"},
		{"run", "reload-environment"},
		{"run", "log", "alfa", "clear"},
	}
}
