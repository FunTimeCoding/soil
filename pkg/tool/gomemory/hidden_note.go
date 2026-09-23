package gomemory

import "fmt"

func hiddenNote(hidden int) string {
	if hidden == 0 {
		return ""
	}

	return fmt.Sprintf(
		"\n%d withheld from this surface; token counts include them\n",
		hidden,
	)
}
