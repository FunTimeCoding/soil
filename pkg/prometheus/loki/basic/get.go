package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Get(l string) string {
	if c.verbose {
		fmt.Printf("GET %s\n", l)
	}

	r := web.NewGet(l)
	r.SetBasicAuth(c.user, c.password)
	s := web.Send(web.Client(), r)

	if c.verbose {
		fmt.Println("Request:")
		fmt.Println(r)
		fmt.Println("Response:")
		fmt.Println(s)

		if s.StatusCode >= 400 {
			fmt.Printf("Error: %s\n", web.ReadString(s))

			return ""
		}
	}

	return web.ReadString(s)
}
