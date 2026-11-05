package unit

func sessionCommands(session string) [][]string {
	return [][]string{
		{"session", "show", session},
		{"session", "tools", session},
		{"session", "cruise", session},
		{"session", "edit", session, "--name", "bravo"},
		{"session", "export", session},
		{"session", "read", session, "kilo"},
		{"session", "peek", session},
		{"session", "print", session},
		{"session", "context", session, "Bash"},
	}
}
