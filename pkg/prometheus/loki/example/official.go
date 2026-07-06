package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Official() {
	host := environment.Required(constant.HostEnvironment)
	user := environment.Required(constant.UserEnvironment)
	password := environment.Required(constant.PasswordEnvironment)

	if false {
		fmt.Printf("Host: %s\n", host)
		fmt.Printf("User: %s\n", user)
		fmt.Printf("Password: %s\n", password)
	}

	// Imports have a big trail of dependencies which would need replaces.
	// 	"github.com/grafana/loki/v3/pkg/logcli/client"
	//	"github.com/grafana/loki/v3/pkg/logproto"
	//c := &client.DefaultClient{
	//	Address:  locator.New(host).String(),
	//	Username: user,
	//	Password: password,
	//	OrgID:    "",
	//}
	//startTime := time.Now()
	//endTime := startTime.Add(-1 * time.Hour)
	//resp, err := c.QueryRange(
	//	`{namespace="bot", source!="event-exporter", stream="stdout"}`,
	//	100,
	//	startTime,
	//	endTime,
	//	logproto.BACKWARD,
	//	time.Minute,
	//	0,
	//	false,
	//)
	//errors.PanicOnError(err)
	//fmt.Printf("Response: %+v\n", resp)
}
