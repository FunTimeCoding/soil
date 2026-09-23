package gomemory

import "fmt"

func withheldNote(withheld int) string {
	if withheld == 0 {
		return ""
	}

	return fmt.Sprintf(
		"%d withheld from this listing; the spreads include them\n",
		withheld,
	)
}
