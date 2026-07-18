package channel

type Channel struct {
	Name       string `gorm:"index;column:name"`
	Identifier int64  `gorm:"primaryKey;autoIncrement:false;column:identifier"`
}
