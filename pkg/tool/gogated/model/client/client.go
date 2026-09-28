package client

import "time"

type Client struct {
	Identifier              string    `gorm:"primaryKey;column:identifier"`
	Secret                  string    `gorm:"column:secret"`
	RedirectLocators        string    `gorm:"column:redirect_locators"`
	GrantTypes              string    `gorm:"column:grant_types"`
	ResponseTypes           string    `gorm:"column:response_types"`
	Scopes                  string    `gorm:"column:scopes"`
	Public                  bool      `gorm:"column:public"`
	TokenEndpointAuthMethod string    `gorm:"column:token_endpoint_auth_method"`
	CreatedAt               time.Time `gorm:"column:created_at"`
}
