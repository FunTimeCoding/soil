package preseed

import "fmt"

func Name(release string) string {
	return fmt.Sprintf("preseed-%s.cfg", release)
}
