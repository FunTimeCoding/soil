package gocredentiald

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"strings"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	o.Database = environment.Required(keepassConstant.DatabaseEnvironment)

	if revealed := environment.Optional(
		constant.RevealedFieldEnvironment,
	); revealed != "" {
		o.RevealedField = strings.Split(revealed, ",")
	}

	if environment.Exists(keepassConstant.PasswordEnvironment) {
		o.Password = environment.Required(keepassConstant.PasswordEnvironment)
	} else {
		o.Password = keychain.Password(constant.Identity.Name())
	}

	Run(o, s)
}
