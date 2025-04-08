package contracts

import (
	"fmt"

	"github.com/cockroachdb/errors/errbase"
)

type Causer interface {
	Cause() error
}

type StackTracer interface {
	StackTrace() errbase.StackTrace
}

type Wrapper interface {
	Unwrap() error
}

type Formatter interface {
	Format(f fmt.State, verb rune)
}

type BaseError interface {
	error
	Wrapper
	Causer
	Formatter
}
