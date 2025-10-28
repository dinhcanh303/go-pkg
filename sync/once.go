package sync

import "sync"

func Once(fn func()) func() {
	return sync.OnceFunc(fn)
}
