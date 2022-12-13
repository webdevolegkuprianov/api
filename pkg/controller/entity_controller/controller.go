package entity_controller

import (
	"net/http"
	"test/pkg/controller"
	"test/pkg/domains/entity"
)

type EntityController struct {
	Domain entity.IDomainEntity
}

func (c *EntityController) HandlerReadEntity(w http.ResponseWriter, r *http.Request) {
	result, err := c.Domain.Read()
	if err != nil {
		controller.ResponseError(w, 500, err)
		return
	}
	controller.Response(w, 200, result)
}
