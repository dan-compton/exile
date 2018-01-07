package plugins

// Caller is a function that can be called.
type Caller interface {
	Call(...interface{}) (string, error)
	Namespace() string
}
