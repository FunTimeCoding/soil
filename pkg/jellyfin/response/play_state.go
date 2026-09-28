package response

type PlayState struct {
	IsPaused      bool  `json:"IsPaused"`
	IsMuted       bool  `json:"IsMuted"`
	PositionTicks int64 `json:"PositionTicks"`
	VolumeLevel   int   `json:"VolumeLevel"`
}
