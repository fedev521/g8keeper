package srv

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
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

var routes = Routes{

	Route{
		"ListPasswords",
		"GET",
		"/v1/passwords",
		ListPasswords,
	},

	Route{
		"PostPassword",
		"POST",
		"/v1/passwords",
		PostPassword,
	},
}
