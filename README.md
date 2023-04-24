# go-joinederror

[![Tests](https://github.com/MrFlynn/go-joinederror/actions/workflows/test.yml/badge.svg)](https://github.com/MrFlynn/go-joinederror/actions/workflows/test.yml)

Small library for unpacking errors into a slice that were combined using the
`errors.Join()` function in Go 1.20.

## Install

```bash
$ go get github.com/mrflynn/go-joinederror
```

## Usage

Below is an example of how to use this library.

```go
// Create a joined error.
err := errors.Join(firstErr, secondErr)

// Unpack and interate.
errs := joinederrors.UnwrapMany(err)
if errs != nil {
    for _, e := range errs {
        fmt.Println(e) // Prints firstErr and then secondErr.
    }
}
```
