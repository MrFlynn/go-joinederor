# go-joinederror

[![Tests](https://github.com/MrFlynn/go-joinederror/actions/workflows/test.yml/badge.svg)](https://github.com/MrFlynn/go-joinederror/actions/workflows/test.yml)

Small library for unpacking errors into a slice that were combined using the
`errors.Join()` function in Go 1.20.

## Install

```bash
$ go get github.com/mrflynn/go-joinederror
```

## Usage

This library provides two ways to unpack joined errors. `UnwrapMany` unpacks
the top level joined errors, and `UnwrapAll` recursively unpacks all joined
errors at every level. For example,

```go
// Create a joined error.
err := errors.Join(
	errors.New("lorem"),
	errors.New("ipsum"),
	errors.Join(errors.New("dolor"), errors.New("sit")),
)

// UnwrapMany:
comparison := []error{
	errors.New("lorem"),
	errors.New("ipsum"),
	errors.Join(errors.New("dolor"), errors.New("sit"))
}

fmt.Println(reflect.DeepEqual(joinederror.UnwrapMany(err), comparison)) // Prints true

// UnwrapAll:
comparison := []error{
	errors.New("lorem"),
	errors.New("ipsum"),
	errors.New("dolor"),
	errors.New("sit"),
}

fmt.Println(reflect.DeepEqual(joinederror.UnwrapAll(err), comparison)) // Prints true
```
