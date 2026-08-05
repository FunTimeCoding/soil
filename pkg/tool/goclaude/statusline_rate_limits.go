package goclaude

type statuslineRateLimits struct {
	FiveHour statuslineRateWindow `json:"five_hour"`
	SevenDay statuslineRateWindow `json:"seven_day"`
}
