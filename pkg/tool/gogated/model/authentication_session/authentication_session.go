package authentication_session

import "time"

type AuthenticationSession struct {
	Identifier      string    `gorm:"primaryKey;column:identifier"`
	UserIdentifier  string    `gorm:"column:user_identifier"`
	UserAgent       string    `gorm:"column:user_agent"`
	Address         string    `gorm:"column:address"`
	AuthenticatedAt time.Time `gorm:"column:authenticated_at"`
	LastUsedAt      time.Time `gorm:"column:last_used_at"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}
