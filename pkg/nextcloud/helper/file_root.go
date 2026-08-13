package helper

import "fmt"

func FileRoot(
	host string,
	user string,
) string {
	return fmt.Sprintf("%s/remote.php/dav/files/%s/", Base(host), user)
}
