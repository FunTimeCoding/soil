package choice

type Choice struct {
	Identifier         uint   `gorm:"primaryKey;autoIncrement;column:identifier"`
	DecisionIdentifier uint   `gorm:"column:decision_identifier;not null;index"`
	Label              string `gorm:"column:label;not null"`
	Position           int    `gorm:"column:position;not null"`
}
