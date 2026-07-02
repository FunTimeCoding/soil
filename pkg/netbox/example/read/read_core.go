package read

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readCore(
	_ *netbox.Client,
	_ *option.Format,
) {
	// Move DataSource here once it works with ExportTemplate (which goes to readExtra)
}
