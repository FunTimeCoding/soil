package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"

func pipInstallCall(c *shell_script.Call) bool {
	pip := c.Name == "pip" || c.Name == "pip3"

	if !pip && c.Name != "python" && c.Name != "python3" {
		return false
	}

	module := false

	for i, argument := range c.Words[1:] {
		literal := argument.Lit()

		if !pip && literal == "-m" && i+2 < len(c.Words) {
			next := c.Words[i+2].Lit()

			if next == "pip" || next == "pip3" {
				module = true
			}
		}

		if literal == "install" && (pip || module) {
			return true
		}
	}

	return false
}
