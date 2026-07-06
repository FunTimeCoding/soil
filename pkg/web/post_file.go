package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log"
)

func PostFile(
	name string,
	locator string,
	user string,
	password string,
) {
	// HTTP upload closes it
	f := system.Open(name)
	r := NewPostBytes(locator, f)
	r.SetBasicAuth(user, password)
	r.Header.Set(constant.ContentType, constant.Form)
	r.ContentLength = system.FileStat(f).Size()
	s := Send(Client(), r)
	defer errors.PanicClose(s.Body)

	if !ResponseOkay(s) {
		log.Panicf("upload %d: %s", s.StatusCode, ReadString(s))
	}
}
