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

	return err.(multiErrorUnwrapper).Unwrap()
}

func prependError(errors []error, err error) []error {
	errors = append(errors, nil)
	copy(errors[1:], errors)
	errors[0] = err

	return errors
}

// UnwrapAll is similar to UnwrapMany except it unwraps joined errors recursively.
func UnwrapAll(err error) []error {
	var (
		allErrors = []error{}
		errStack  = newStack[error]()
	)

	errStack.push(err)

	for !errStack.empty() {
		top, _ := errStack.pop()

		if top == nil {
			continue
		}

		if !implementsMultiErrorUnwrapper(top) {
			allErrors = prependError(allErrors, top)
		} else {
			for _, e := range top.(multiErrorUnwrapper).Unwrap() {
				errStack.push(e)
			}
		}
	}

	return allErrors
}
