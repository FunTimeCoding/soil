package attribution

import "time"

type Attribution struct {
	PlaceKind       string    `gorm:"column:place_kind"`
	PlaceIdentifier int32     `gorm:"column:place_identifier"`
	PlaceName       string    `gorm:"column:place_name"`
	SeenAt          time.Time `gorm:"column:seen_at"`
}
