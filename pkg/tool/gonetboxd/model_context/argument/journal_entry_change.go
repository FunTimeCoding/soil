package argument

type JournalEntryChange struct {
	Identifier float64 `json:"identifier"`
	Comments   string  `json:"comments"`
	Kind       string  `json:"kind"`
}
