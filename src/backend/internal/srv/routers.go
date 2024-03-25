package srv

import (
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/store"
	"github.com/fedev521/g8keeper/backend/internal/svc"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(tinkSvcConf svc.TinkSvcConfig) *mux.Router {
	var keeper = store.NewInMapKeeper()

	var routes = Routes{
		Route{
			"ListPasswords",
			"GET",
			"/v1/passwords",
			ListPasswordsHF(keeper),
		},
		Route{
			"PostPassword",
			"POST",
			"/v1/passwords",
			PostPasswordHF(keeper),
		},
	}

	router := mux.NewRouter().StrictSlash(true)
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
