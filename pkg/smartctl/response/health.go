package response

type Health struct {
	CriticalWarning int   `json:"critical_warning"`
	AvailableSpare  int   `json:"available_spare"`
	PercentageUsed  int   `json:"percentage_used"`
	MediaErrors     int64 `json:"media_errors"`
}
