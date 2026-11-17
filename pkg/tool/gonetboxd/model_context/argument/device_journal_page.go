package argument

type DeviceJournalPage struct {
	Device string  `json:"device"`
	Limit  float64 `json:"limit"`
	Offset float64 `json:"offset"`
}
