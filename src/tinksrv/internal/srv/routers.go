package srv

import (
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/internal/kms"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(kekManager *kms.KEKManager) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"PostEncrypt",
			"POST",
			"/v1/encrypt",
			PostEncryptHF(kekManager),
		},

		Route{
			"PostDecrypt",
			"POST",
			"/v1/decrypt",
			PostDecryptHF(kekManager),
		},
	}

	for _, r := range routes {
		var handler http.Handler
		handler = r.HandlerFunc
		handler = Logger(handler, r.Name)

		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(handler)
	}

	return router
}
