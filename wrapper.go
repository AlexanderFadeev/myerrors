package myerrors

type Wrapper interface {
	Unwrap() error
}

type wrapper struct {
	error
}

type causer interface {
	Cause() error
}

func (w *wrapper) Unwrap() error {
	res := w.error.(causer).Cause()
	// res is always != nil

	if _, ok := res.(Wrapper); ok {
		return res
	}

	if _, ok := res.(causer); ok {
		return &wrapper{res}
	}

	return res
}
