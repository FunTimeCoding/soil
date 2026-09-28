package user

import "time"

type User struct {
	Identifier          string    `gorm:"primaryKey;column:identifier"`
	Mail                string    `gorm:"uniqueIndex;column:mail"`
	PasswordHash        string    `gorm:"column:password_hash"`
	Source              string    `gorm:"column:source;default:local"`
	DirectoryIdentifier *string   `gorm:"uniqueIndex;column:directory_identifier"`
	Account             string    `gorm:"column:account"`
	Name                string    `gorm:"column:name"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}
