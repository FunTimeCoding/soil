package installed

type Installed struct {
	Identifier       uint   `gorm:"primarykey;column:id"`
	SystemIdentifier uint   `gorm:"not null;index"`
	Name             string `gorm:"not null"`
	Version          string `gorm:"not null"`
	Source           string `gorm:"not null"`
	SourceVersion    string `gorm:"not null"`
	Candidate        string `gorm:"not null;default:''"`
}
