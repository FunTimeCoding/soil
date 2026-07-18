package label

type Label struct {
	Identifier       uint   `gorm:"primaryKey;autoIncrement;column:identifier"`
	ObjectType       string `gorm:"uniqueIndex:object_key;column:object_type"`
	ObjectIdentifier int32  `gorm:"uniqueIndex:object_key;column:object_identifier"`
	Key              string `gorm:"uniqueIndex:object_key;column:key"`
	Value            string `gorm:"column:value"`
}
