package repository

type Cache interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (value string, err error)
	Delete(key string) (err error)
	DeleteAll() (err error)
	Shutdown() (err error)
}
