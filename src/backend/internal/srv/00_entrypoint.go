package srv

import (
	"fmt"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
	"github.com/fedev521/g8keeper/backend/internal/svc"
)

func StartServer(conf Config, tinkSvcConf svc.TinkSvcConfig) error {
	router := NewRouter(tinkSvcConf)
	log.Info("Starting server", map[string]interface{}{
		"name":    conf.Name,
		"port":    conf.Port,
		"tinksvc": tinkSvcConf,
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router)
}
