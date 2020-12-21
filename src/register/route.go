package register

import (
	"github.com/Etpmls/EM-Auth/src/application/service"
	em_library "github.com/Etpmls/Etpmls-Micro/library"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

// Register Route
func RegisterRoute(mux *runtime.ServeMux)  {
	mux.HandlePath("GET", em_library.Config.ServiceDiscovery.Service.CheckUrl, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello"))
	})
	mux.HandlePath("GET", "/api/checkAuth", service.ServiceAuth{}.Check)
}