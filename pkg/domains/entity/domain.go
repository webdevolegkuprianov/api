package entity

import (
	"context"

	"github.com/rs/zerolog"
)

type Domain struct {
	l         zerolog.Logger
	daoEntity iEntityDAO
}

func NewDomain(ctx context.Context) *Domain {
	l := zerolog.Ctx(ctx).With().
		Str("Package", "entity").
		Str("Class", "Domain").
		Logger()

	d := &Domain{
		l:         l,
		daoEntity: newEntityDAO(ctx),
	}
	return d
}

func (d *Domain) Read() (result *EntityDTO, err error) {
	result, err = d.daoEntity.read()
	if err != nil {
		return &EntityDTO{}, err
	}
	return
}
