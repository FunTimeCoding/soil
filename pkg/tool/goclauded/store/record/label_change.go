package record

type LabelChange struct {
	EventIdentifier uint   `gorm:"column:event_identifier"`
	Value           string `gorm:"column:value"`
}
