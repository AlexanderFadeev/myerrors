package myerrors

import "sync"

type Handler interface {
	Handle(error)
}

func SafeHandler(handler Handler) Handler {
	return &safeHandler{
		impl: handler,
	}
}

type safeHandler struct {
	impl Handler
	mu   sync.Mutex
}

func (a *safeHandler) Handle(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.impl.Handle(err)
}
