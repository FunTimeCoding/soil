package constant

const (
	HostEnvironment        = "LDAP_HOST"
	PortEnvironment        = "LDAP_PORT"
	InsecureEnvironment    = "LDAP_INSECURE"
	UntrustedEnvironment   = "LDAP_UNTRUSTED"
	BindEnvironment        = "LDAP_BIND_DISTINGUISHED_NAME"
	PasswordEnvironment    = "LDAP_BIND_PASSWORD"
	BaseEnvironment        = "LDAP_BASE_DISTINGUISHED_NAME"
	UserFilterEnvironment  = "LDAP_USER_FILTER"
	GroupFilterEnvironment = "LDAP_GROUP_FILTER"
	AuthorityEnvironment   = "LDAP_AUTHORITY_FILE"

	SecurePort   = 636
	InsecurePort = 389

	UniqueAttribute  = "entryUUID"
	MailAttribute    = "mail"
	NameAttribute    = "cn"
	AccountAttribute = "uid"

	Subject = "directory"
)
