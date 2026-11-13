package opening_row

func New(
	identifier string,
	open string,
	close string,
) *OpeningRow {
	return &OpeningRow{identifier: identifier, open: open, close: close}
}
