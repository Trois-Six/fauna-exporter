// Package handlers contains all the HTTP handlers.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/prometheus/common/version"
)

// Handlers a set of handlers.
type Handlers struct {
	metricsPath string
}

// New creates all HTTP handlers.
func New(metricsPath string) Handlers {
	return Handlers{
		metricsPath: metricsPath,
	}
}

// Index is the landing page handler.
func (h Handlers) Index(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, `<html>
	<head><title>Fauna Exporter</title></head>
	<body>
	<h1>Fauna Exporter</h1>
	<p><a href='`+h.metricsPath+`'>Metrics</a></p>
	<h2>Build</h2>
	<pre>`+version.Info()+` `+version.BuildContext()+`</pre>
	</body>
	</html>`)
}

// OK is the healthCheck handler.
func (h Handlers) OK(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "OK")
}
