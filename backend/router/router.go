package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(AccessLogMiddleware)
	router.Handle("/stats/prometheus", promhttp.Handler()).Methods(http.MethodGet)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/liveness-probe", LivenessProbeHandler).Methods(http.MethodGet)
	api.HandleFunc("/request", RequestHandler).Methods(http.MethodGet)
	api.Use(RequestCountingMiddleware)

	admin := router.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/overload", OverloadHandler).Methods(http.MethodGet)
	admin.HandleFunc("/error", ErrorHandler).Methods(http.MethodGet)
	admin.HandleFunc("/reboot", RebootHandler).Methods(http.MethodGet)

	return router
}