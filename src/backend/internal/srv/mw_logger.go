package srv

import (
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
)

func Logger(next http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Received request", map[string]interface{}{
			"method": r.Method,
			"uri":    r.RequestURI,
			"route":  routeName,
		})
		next.ServeHTTP(w, r)

		// TODO add id to the request and log time needed for processing
	})
}
