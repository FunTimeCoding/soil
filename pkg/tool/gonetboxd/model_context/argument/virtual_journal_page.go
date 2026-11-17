package argument

type VirtualJournalPage struct {
	VirtualMachine string  `json:"virtual_machine"`
	Limit          float64 `json:"limit"`
	Offset         float64 `json:"offset"`
}
