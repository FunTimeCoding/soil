package argument

type DeviceJournalEntry struct {
	Device   string `json:"device"`
	Comments string `json:"comments"`
	Kind     string `json:"kind"`
}
