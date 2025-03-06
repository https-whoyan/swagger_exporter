package sheets

import "net/http"

type Client interface {
	SafeToList(httpCli *http.Client, bu
}
