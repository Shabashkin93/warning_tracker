package transport

type Status interface {
	Register(string, interface{})
}
