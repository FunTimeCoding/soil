package opening_row

import "maragu.dev/gomponents"

type OpeningRow struct {
	identifier string
	open       string
	close      string
	cells      []gomponents.Node
	columns    int
	detail     []gomponents.Node
	expanded   bool
}
