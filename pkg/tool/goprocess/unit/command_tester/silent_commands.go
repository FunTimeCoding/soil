package command_tester

func SilentCommands() [][]string {
	return [][]string{
		{"run", "start", "bravo"},
		{"run", "stop", "alfa"},
		{"run", "restart", "alfa"},
		{"run", "reload-procfile"},
		{"run", "reload-environment"},
		{"run", "log", "alfa", "clear"},
	}
}
