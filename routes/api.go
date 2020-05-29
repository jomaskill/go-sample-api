package api

import (
	actionUseCase "TesteWeb/application"
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/action", actionUseCase.Create).Methods("POST")
	router.HandleFunc("/action", actionUseCase.Index).Methods("GET")
	router.HandleFunc("/action/{uuid}", actionUseCase.Show).Methods("GET")
	router.HandleFunc("/action/{uuid}", actionUseCase.Update).Methods("PATCH", "PUT")
	router.HandleFunc("/action/{uuid}", actionUseCase.Destroy).Methods("DELETE")
}