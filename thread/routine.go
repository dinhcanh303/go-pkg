package thread

import (
	"context"
	"go-pkg/rec"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	go RunSafe(fn)
}

// GoSafeCtx runs the given fn using another goroutine, recovers if fn panics with ctx.
func GoSafeCtx(ctx context.Context, fn func()) {
	go RunSafeCtx(ctx, fn)
}

// RunSafe runs the given fn, recovers if fn panics.
func RunSafe(fn func()) {
	defer rec.Recover()

	fn()
}

// RunSafeCtx runs the given fn, recovers if fn panics with ctx.
func RunSafeCtx(ctx context.Context, fn func()) {
	// defer rec.RecoverCtx(ctx)

	fn()
}
