package netboot

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Link(
	release string,
	architecture string,
) string {
	return locator.New("ftp.debian.org").Insecure().Path(
		fmt.Sprintf(
			"debian/dists/%s/main/installer-%s/current/images/netboot/netboot.tar.gz",
			release,
			architecture,
		),
	).String()
}
