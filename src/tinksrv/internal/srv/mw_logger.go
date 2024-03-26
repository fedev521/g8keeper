package srv

import (
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/internal/log"
)

func Logger(inner http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Received request", map[string]interface{}{
			"method": r.Method,
			"uri":    r.RequestURI,
			"route":  routeName,
		})
		inner.ServeHTTP(w, r)

		// TODO add id to the request and log time needed for processing
	})
}
