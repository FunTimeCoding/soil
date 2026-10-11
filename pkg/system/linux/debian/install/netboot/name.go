package netboot

import "fmt"

func Name(release string) string {
	return fmt.Sprintf("netboot-%s.tar.gz", release)
}
