package goclaude

type StatusLineRateLimits struct {
	FiveHour    *StatusLineRateWindow   `json:"five_hour"`
	SevenDay    *StatusLineRateWindow   `json:"seven_day"`
	ModelScoped []StatusLineModelScoped `json:"model_scoped"`
}
