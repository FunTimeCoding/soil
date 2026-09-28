package fritz

import (
	"fmt"
	"net/http"
)

func soapHeader(
	request *http.Request,
	service string,
	action string,
) {
	request.Header.Set("Content-Type", `text/xml; charset="utf-8"`)
	request.Header.Set("SOAPACTION", fmt.Sprintf(`"%s#%s"`, service, action))
}
