package claude_tester

import "fmt"

func UserLine(text string) string {
	return fmt.Sprintf(`{"type":"user","message":{"content":"%s"}}`, text)
}
