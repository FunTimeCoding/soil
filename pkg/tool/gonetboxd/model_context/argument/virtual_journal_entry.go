package argument

type VirtualJournalEntry struct {
	VirtualMachine string `json:"virtual_machine"`
	Comments       string `json:"comments"`
	Kind           string `json:"kind"`
}
