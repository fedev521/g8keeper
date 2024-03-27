package srv

import (
	"fmt"
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/internal/kms"
	"github.com/fedev521/g8keeper/tinksrv/internal/log"
)

func StartServer(conf Config, kekManager *kms.KEKManager) error {
	router := NewRouter(kekManager)
	log.Info("Starting server", map[string]interface{}{
		"name": conf.Name,
		"port": conf.Port,
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router)
}
