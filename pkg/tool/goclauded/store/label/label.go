package label

type Label struct {
	Identifier        uint   `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string `gorm:"uniqueIndex:session_key_value;column:session_identifier"`
	Key               string `gorm:"uniqueIndex:session_key_value;column:key"`
	Value             string `gorm:"uniqueIndex:session_key_value;column:value"`
}
