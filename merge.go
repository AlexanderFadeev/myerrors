package myerrors

func Merge(errA, errB error) error {
	switch {
	case errA != nil:
		return errA
	case errB != nil:
		return errB
	default:
		return nil
	}
}
