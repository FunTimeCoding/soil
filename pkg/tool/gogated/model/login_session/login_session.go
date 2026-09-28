package login_session

import "time"

type LoginSession struct {
	Identifier       string    `gorm:"primaryKey;column:identifier"`
	AuthorizeRequest string    `gorm:"column:authorize_request"`
	UserIdentifier   string    `gorm:"column:user_identifier"`
	Authenticated    bool      `gorm:"column:authenticated"`
	CreatedAt        time.Time `gorm:"column:created_at"`
}
