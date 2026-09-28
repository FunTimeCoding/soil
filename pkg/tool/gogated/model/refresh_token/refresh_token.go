package refresh_token

import "time"

type RefreshToken struct {
	Signature         string    `gorm:"primaryKey;column:signature"`
	RequestIdentifier string    `gorm:"column:request_identifier"`
	ClientIdentifier  string    `gorm:"column:client_identifier"`
	Scopes            string    `gorm:"column:scopes"`
	GrantedScopes     string    `gorm:"column:granted_scopes"`
	Session           string    `gorm:"column:session_data"`
	Form              string    `gorm:"column:form_data"`
	Active            bool      `gorm:"column:active"`
	RequestedAt       time.Time `gorm:"column:requested_at"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
