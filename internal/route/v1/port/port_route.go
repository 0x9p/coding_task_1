package port

import "github.com/gorilla/mux"

func Route(router *mux.Router, handler *Handler) {
	subRouter := router.PathPrefix("/ports").Subrouter()

	subRouter.HandleFunc("/batch", handler.HandleUpsertPorts()).Methods("POST")
}
