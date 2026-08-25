package response

import "time"

type User struct {
	Identifier      string     `json:"id"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	AvatarLocator   string     `json:"avatarUrl"`
	IsActive        bool       `json:"isActive"`
	HasPasswordAuth bool       `json:"hasPasswordAuth"`
	IsManaged       bool       `json:"isManaged"`
	IsSuperuser     bool       `json:"isSuperuser"`
	IsStaff         bool       `json:"isStaff"`
	Has2FA          bool       `json:"has2fa"`
	DateJoined      *time.Time `json:"dateJoined"`
	LastLogin       *time.Time `json:"lastLogin"`
	LastActive      *time.Time `json:"lastActive"`
	Permissions     []string   `json:"permissions"`
}
