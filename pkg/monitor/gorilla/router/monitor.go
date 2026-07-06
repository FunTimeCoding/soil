package router

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/flag"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"log"
	"net"
	"net/http"
)

func (r *Router) Monitor(
	w http.ResponseWriter,
	q *http.Request,
) {
	c, upgradeFail := constant.Upgrader.Upgrade(w, q, nil)

	if upgradeFail != nil {
		log.Printf("upgrade: %s\n", upgradeFail)

		return
	}

	defer errors.LogClose(c)

	for {
		messageType, text, e := c.ReadMessage()

		if e != nil {
			log.Printf("read: %s\n", e)

			break
		}

		arguments := split.Comma(string(text))
		address := net.ParseIP(network.SplitHost(q.RemoteAddr))
		log.Printf("%s: %+v\n", address, arguments)
		i := r.FindClient(address)

		if i == nil {
			log.Printf("client not found: %s\n", address)

			continue
		}

		switch arguments[0] {
		case constant.LoginCommand:
			i = client.New(arguments[1], arguments[2], address, c)
			r.Client = append(r.Client, i)
			i.Send(
				join.Comma(
					[]string{constant.LoginResponseCommand, address.String()},
				),
			)

			for _, l := range r.Client {
				log.Printf("client: %+v\n", l)
			}
		case constant.LogoutCommand:
			for j, l := range r.Client {
				if l.Address.Equal(address) {
					r.Client = append(r.Client[:j], r.Client[j+1:]...)
				} else {
					log.Printf("client: %+v\n", l)
				}
			}
		case constant.FlagCommand:
			identifier := arguments[1]
			var operation string

			if r.HasFlag(identifier) {
				l := r.FlagByIdentifier(identifier)

				if l == nil {
					log.Printf("flag not found: %s\n", identifier)

					continue
				}

				if l.HasClient(i) {
					operation = constant.FlagRemoveCommand
					l.RemoveClient(i)

					if len(l.By) == 0 {
						r.RemoveFlag(l)
					}
				} else {
					operation = constant.FlagAddCommand
					l.AddClient(i)
				}
			} else {
				operation = constant.FlagAddCommand
				r.AddFlag(flag.New(identifier, i))
			}

			for _, l := range r.Client {
				l.Send(
					join.Comma(
						[]string{operation, i.Handle(), arguments[1]},
					),
				)
			}
		case constant.ClearCommand:
			var found bool

			for _, l := range r.Flag {
				for j, k := range l.By {
					if k.Address.Equal(i.Address) {
						if len(l.By) == 1 {
							r.Flag = append(r.Flag[:j], r.Flag[j+1:]...)
						} else {
							l.By = append(l.By[:j], l.By[j+1:]...)
						}

						found = true

						break
					}
				}

				if found {
					break
				}
			}

			for _, l := range r.Client {
				if l.Address.Equal(i.Address) {
					continue
				}

				l.Send(
					join.Comma(
						[]string{
							constant.FlagRemoveCommand,
							i.Handle(),
							arguments[1],
						},
					),
				)
			}
		case constant.PingCommand:
		default:
			if f := c.WriteMessage(messageType, text); f != nil {
				log.Printf("write: %s\n", f)
			}
		}
	}
}
