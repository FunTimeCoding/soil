package gomemory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func spreadRow(
	name string,
	s client.Spread,
) string {
	return fmt.Sprintf(
		"%-12s %7d %9d %12d %8d",
		name,
		s.Median,
		s.Ninetieth,
		s.NinetyNinth,
		s.Maximum,
	)
}
