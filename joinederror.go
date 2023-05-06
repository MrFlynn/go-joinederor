package joinederror

import "reflect"

type multiErrorUnwrapper interface {
	Unwrap() []error
}

var impl = reflect.TypeOf((*multiErrorUnwrapper)(nil)).Elem()

func implementsMultiErrorUnwrapper(err error) bool {
	return reflect.TypeOf(err).Implements(impl)
}

// UnwrapMany takes a joined error from errors.Join and returns the slice of
// errors that make up the joined error. Returns nil if the supplied error is
// nil or is not a joined error.
func UnwrapMany(err error) []error {
	if err == nil {
		return nil
	}

	if !implementsMultiErrorUnwrapper(err) {
		return nil
	}

	errors := err.(multiErrorUnwrapper)
	return errors.Unwrap()
}
