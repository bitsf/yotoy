package pkg

var factory = make(map[string]Handler)

func Register(name string, handler Handler) {
	factory[name] = handler
}

func GetFactory() map[string]Handler {
	return factory
}
