package entity

import (
	"context"
	"fmt"
)

type property string

const (
	property1 property = "property"
)

type entityDAO struct {
	ctx     context.Context
	storage map[property]string
}

func newEntityDAO(ctx context.Context) *entityDAO {
	mp := map[property]string{}
	mp[property1] = "property"
	return &entityDAO{
		ctx:     ctx,
		storage: mp,
	}
}

func (dao *entityDAO) read() (entity *EntityDTO, err error) {
	prop, ok := dao.storage[property1]
	if !ok {
		return &EntityDTO{}, fmt.Errorf("ошибка получения property")
	}
	entity = &EntityDTO{
		Property1: prop,
	}
	return
}
