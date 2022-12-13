package router

import (
	"test/pkg/controller/entity_controller"

	"github.com/gorilla/mux"
)

type Router struct {
	entityController entity_controller.EntityController
	Router           *mux.Router
}

func NewRouter(entityController *entity_controller.EntityController) *Router {
	return &Router{
		entityController: *entityController,
		Router:           mux.NewRouter().StrictSlash(true),
	}
}

func (r *Router) RouterEntityInit() {
	routerApiV1 := r.Router.PathPrefix("/api/v1").Subrouter()
	routerApiV1.HandleFunc("/checkcar", r.entityController.HandlerReadEntity).Methods("GET", "OPTIONS")
}
