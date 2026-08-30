package response

type Event struct {
	Tool      string `json:"tool"`
	Surface   string `json:"surface"`
	Actor     string `json:"actor"`
	Outcome   string `json:"outcome"`
	Kind      string `json:"kind"`
	Duration  int64  `json:"duration_ms,omitzero"`
	CreatedAt string `json:"created_at"`
}
