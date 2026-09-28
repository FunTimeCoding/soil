package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gogated",
	"OAuth 2.0 / OIDC authorization server",
	"gogated",
)

const (
	SecretEnvironment                = "GATE_SECRET"
	SuperUserMailEnvironment         = "GATE_SUPER_USER_MAIL"
	SuperUserPasswordEnvironment     = "GATE_SUPER_USER_PASSWORD"
	IssuerEnvironment                = "GATE_ISSUER"
	AdminClientIdentifierEnvironment = "GATE_ADMIN_CLIENT_IDENTIFIER"
	AdminClientSecretEnvironment     = "GATE_ADMIN_CLIENT_SECRET"

	CookieName               = "gogated_login"
	AuthenticationCookieName = "gogated_session"

	MailField         = "mail"
	PasswordField     = "password"
	PasswordInputType = "password"

	ClientKind   = "client"
	SecretNotice = "the secret is shown only once - save it now"

	ListClient   = "list_client"
	CreateClient = "create_client"
	UpdateClient = "update_client"
	DeleteClient = "delete_client"

	SourceField     = "source"
	SourceLocal     = "local"
	SourceDirectory = "directory"

	LocalLabel               = "Local account"
	DirectoryLabel           = "Directory"
	MailLabel                = "Mail"
	AccountLabel             = "Mail or account"
	LocalRejectedMessage     = "No local account with those credentials"
	DirectoryRejectedMessage = "The directory rejected those credentials"

	LoginSessionTimeToLive = 5 * time.Minute

	AuthenticationIdleTimeToLive     = 7 * 24 * time.Hour
	AuthenticationAbsoluteTimeToLive = 30 * 24 * time.Hour

	ClientsTitle             = "Clients"
	ClientsPath              = "/"
	CreateTitle              = "Create Client"
	CreatePath               = "/clients/create"
	SessionsTitle            = "Sessions"
	SessionsPath             = "/sessions"
	SessionsRevokeOthersPath = "/sessions/revoke-others"
	LogoutPath               = "/logout"
	SignOutTitle             = "Sign out"
	WarningClass             = "warning"

	GrantAuthorizationCode      = "authorization_code"
	GrantRefreshToken           = "refresh_token"
	ResponseCode                = "code"
	PromptParameter             = "prompt"
	PromptLogin                 = "login"
	MaxAgeParameter             = "max_age"
	ScopeOpenIdentity           = "openid"
	ScopeOffline                = "offline"
	AuthMethodClientSecretPost  = "client_secret_post"
	AuthMethodClientSecretBasic = "client_secret_basic"
	AuthMethodNone              = "none"
	AlgorithmRS256              = "RS256"
	KeyTypeRSA                  = "RSA"
	KeyUseSigning               = "sig"
	SubjectTypePublic           = "public"
	ChallengeMethodS256         = "S256"
)

const InlineCSS = `
.credential-card { background: var(--pico-card-background-color); border: 1px solid var(--pico-primary); border-radius: var(--pico-border-radius); padding: 1.5rem; margin: 1rem 0; }
.credential-card code { word-break: break-all; }
.credential-card p { margin-bottom: 0.5rem; }
.warning { color: var(--pico-del-color); font-weight: bold; }
`

const (
	FixtureSecret   = "test-secret-at-least-32-bytes!xy"
	FixtureIssuer   = "http://localhost"
	FixtureMail     = "admin@test.internal"
	FixturePassword = "password"
)
