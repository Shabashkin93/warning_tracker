package repository

type Status interface {
	GetStatus() (status bool, err error)
}
