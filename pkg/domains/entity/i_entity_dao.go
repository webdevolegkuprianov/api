package entity

type iEntityDAO interface {
	read() (entity *EntityDTO, err error)
}
