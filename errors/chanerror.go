package errors

import (
	"errors"
	"fmt"
	"os"
)

func AssertChanErrorWithoutAs(err error) {
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", e.Op, e.Path, e.Err)
	}
}

func AssertChanErrorWithAs(err error) {
	var e *os.PathError
	if errors.As(err, &e) {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", e.Op, e.Path, e.Err)
	}
}
