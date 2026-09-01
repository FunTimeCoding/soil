package service

import (
	"fmt"
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
	"github.com/tobischo/gokeepasslib/v3"
	"time"
)

func describe(
	path string,
	entry *gokeepasslib.Entry,
) *credential.Credential {
	modified := time.Time{}

	if entry.Times.LastModificationTime != nil {
		modified = entry.Times.LastModificationTime.Time
	}

	return credential.New(
		fmt.Sprintf("%x", entry.UUID),
		path,
		entry.GetTitle(),
		entry.GetContent(keepassConstant.UserNameKey),
		entry.GetContent(constant.LocatorKey),
		modified,
	)
}
