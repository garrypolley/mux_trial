package middleware

import (
	"net/http"
	"time"

	"github.com/garrypolley/mux_trial/logging"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
)

var log *logrus.Logger

func init() {
	log = logging.Logger
}

type LoggingMiddleware struct {
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (l LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	log.Infof("Started %s %s", r.Method, r.URL.Path)

	next(w, r)

	res := w.(negroni.ResponseWriter)
	log.Infof("Completed %v %s in %v", res.Status(), http.StatusText(res.Status()), time.Since(start))
}
