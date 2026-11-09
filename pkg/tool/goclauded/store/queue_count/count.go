package queue_count

type Count struct {
	SessionIdentifier string `gorm:"column:session_identifier"`
	Callsign          string `gorm:"column:callsign"`
	Waiting           int64  `gorm:"column:waiting"`
}
