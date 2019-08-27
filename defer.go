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

func CallWrapd(err *error, fn func() error, msg string) {
	callErr := fn()
	callErr = Wrap(callErr, msg)
	*err = Merge(*err, callErr)
}

func CallWrapfd(err *error, fn func() error, format string, args ...interface{}) {
	callErr := fn()
	callErr = Wrapf(callErr, format, args...)
	*err = Merge(*err, callErr)
}
