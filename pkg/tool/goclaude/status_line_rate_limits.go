package goclaude

type statusLineRateLimits struct {
	FiveHour    *statusLineRateWindow   `json:"five_hour"`
	SevenDay    *statusLineRateWindow   `json:"seven_day"`
	ModelScoped []statusLineModelScoped `json:"model_scoped"`
}
