package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/access_token"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authorization_code"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/open_identity_session"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/proof_key"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/refresh_token"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/signing_key"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	errors.PanicOnError(d.AutoMigrate(client.Stub()))
	errors.PanicOnError(d.AutoMigrate(authorization_code.Stub()))
	errors.PanicOnError(d.AutoMigrate(access_token.Stub()))
	errors.PanicOnError(d.AutoMigrate(refresh_token.Stub()))
	errors.PanicOnError(d.AutoMigrate(proof_key.Stub()))
	errors.PanicOnError(d.AutoMigrate(open_identity_session.Stub()))
	errors.PanicOnError(d.AutoMigrate(login_session.Stub()))
	errors.PanicOnError(d.AutoMigrate(authentication_session.Stub()))
	errors.PanicOnError(d.AutoMigrate(signing_key.Stub()))
	errors.PanicOnError(d.AutoMigrate(user.Stub()))
}
