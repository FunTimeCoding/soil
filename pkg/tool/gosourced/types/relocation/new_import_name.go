package relocation

func NewImportName(
	local string,
	imported bool,
) *ImportName {
	return &ImportName{Local: local, Imported: imported}
}
