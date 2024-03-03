package srv

import (
	"fmt"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from demo Tink encryption/decryption server!\n")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"PostEncrypt",
		"POST",
		"/v1/encrypt",
		PostEncrypt,
	},

	Route{
		"PostDecrypt",
		"POST",
		"/v1/decrypt",
		PostDecrypt,
	},
}
