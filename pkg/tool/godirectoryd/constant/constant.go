package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"godirectoryd",
	"LDAP directory user and group management",
	"godirectoryd",
)

const (
	UserContainer  = "ou=people"
	GroupContainer = "ou=groups"

	PersonClass = "inetOrgPerson"
	GroupClass  = "posixGroup"

	AccountAttribute     = "uid"
	NameAttribute        = "cn"
	SurnameAttribute     = "sn"
	MailAttribute        = "mail"
	PasswordAttribute    = "userPassword"
	ClassAttribute       = "objectClass"
	UniqueAttribute      = "entryUUID"
	MemberAttribute      = "memberUid"
	GroupNumberAttribute = "gidNumber"

	PersonFilter = "(objectClass=inetOrgPerson)"
	GroupFilter  = "(objectClass=posixGroup)"

	UserKind  = "directory user"
	GroupKind = "directory group"

	HostEnvironment     = "GODIRECTORY_HOST"
	PortEnvironment     = "GODIRECTORY_PORT"
	InsecureEnvironment = "GODIRECTORY_INSECURE"
	TokenEnvironment    = "GODIRECTORY_TOKEN"

	FirstGroupNumber = 5000

	ListUser     = "list_user"
	CreateUser   = "create_user"
	ModifyUser   = "modify_user"
	DeleteUser   = "delete_user"
	SetPassword  = "set_password"
	ListGroup    = "list_group"
	CreateGroup  = "create_group"
	DeleteGroup  = "delete_group"
	AddMember    = "add_member"
	RemoveMember = "remove_member"

	UserTitle  = "Users"
	UserPath   = "/"
	GroupTitle = "Groups"
	GroupPath  = "/group"

	CreateUserPath   = "/user/create"
	DeleteUserPath   = "/user/delete"
	SetPasswordPath  = "/user/password"
	CreateGroupPath  = "/group/create"
	DeleteGroupPath  = "/group/delete"
	AddMemberPath    = "/group/member/add"
	RemoveMemberPath = "/group/member/remove"

	TextInputType     = "text"
	MailInputType     = "email"
	PasswordInputType = "password"

	AccountField  = "account"
	NameField     = "name"
	SurnameField  = "surname"
	MailField     = "mail"
	PasswordField = "password"
	GroupField    = "group"
)
