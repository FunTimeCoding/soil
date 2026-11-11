package response

type Report struct {
	Device      *Device      `json:"device"`
	ModelName   string       `json:"model_name"`
	Status      *Status      `json:"smart_status"`
	Temperature *Temperature `json:"temperature"`
	Health      *Health      `json:"nvme_smart_health_information_log"`
	Run         *Run         `json:"smartctl"`
}
