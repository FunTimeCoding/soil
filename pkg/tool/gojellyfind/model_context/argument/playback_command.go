package argument

type PlaybackCommand struct {
	SessionIdentifier string `json:"session_id"`
	Command           string `json:"command"`
	SeekPositionTicks int64  `json:"seek_position_ticks"`
}
