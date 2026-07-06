package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(locator.New(c.host).Port(c.port).Path(path).String())
	r.SetBasicAuth(c.user, c.password)
	r.Header.Add(constant.ContentType, constant.Object)
	r.Header.Add(constant.Accept, constant.Object)
	response := web.Send(web.Client(), r)

	if false {
		fmt.Println(r)
		fmt.Println(response)
	}

	return web.ReadString(response)
}
