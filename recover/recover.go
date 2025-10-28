package recover

import "github.com/pkg/errors"

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if p := recover(); p != nil {
		errors.WithStack(p.(error))
	}
}
