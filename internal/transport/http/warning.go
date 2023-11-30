package transport

type Warning interface {
	Register(string, interface{})
}
