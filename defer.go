package myerrors

func Wrapd(err *error, msg string) {
	*err = Wrap(*err, msg)
}

func Wrapfd(err *error, format string, args ...interface{}) {
	*err = Wrapf(*err, format, args...)
}

func Calld(err *error, fn func() error) {
	callErr := fn()
	*err = Merge(*err, callErr)
}
