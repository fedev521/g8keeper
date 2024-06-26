package srv

import (
	"fmt"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
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
	var keeper store.PasswordKeeper
	keeper, err := store.NewCryptedInMemKeeper(tinkSvcConf)
	if err != nil {
		log.Error(err.Error())
		// TODO return err
	}

	var routes = Routes{
		Route{
			"GetPassword",
			"GET",
			fmt.Sprintf("/v1/passwords/{%s}", passwordIdKey),
			GetPasswordHF(keeper),
		},
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
