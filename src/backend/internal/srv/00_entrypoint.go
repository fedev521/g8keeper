package srv

import (
	"fmt"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
)

func StartServer(conf Config) error {
	router := NewRouter()
	log.Info("Starting server", map[string]interface{}{
		"name": conf.Name,
		"port": conf.Port,
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router)
}
