package argument

type Play struct {
	SessionIdentifier  string   `json:"session_id"`
	ItemIDs            []string `json:"item_ids"`
	PlayCommand        string   `json:"play_command"`
	StartPositionTicks int64    `json:"start_position_ticks"`
}
