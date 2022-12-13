package entity

type IDomainEntity interface {
	Read() (result *EntityDTO, err error)
}
