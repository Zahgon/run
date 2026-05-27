package run

import (
	"context"
	"errors"
	"os"
)

// ContextHandler returns an actor, i.e. an execute and interrupt func, that
// terminates when the provided context is canceled.
func ContextHandler(ctx context.Context) (execute func() error, interrupt func(error)) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignalHandler returns an actor, i.e. an execute and interrupt func, that
// terminates with ErrSignal when the process receives one of the provided
// signals, or with ctx.Error() when the parent context is canceled. If no
// signals are provided, the actor will terminate on any signal, per
// [signal.Notify].
func SignalHandler(ctx context.Context, signals ...os.Signal) (execute func() error, interrupt func(error)) {
	_ = "STUB: not implemented"
	return nil, nil
}

type testSigChanKey struct{}

func getTestSigChan(ctx context.Context) <-chan os.Signal { _ = "STUB: not implemented"; return nil }

// can be nil

func putTestSigChan(ctx context.Context, c <-chan os.Signal) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SignalError is returned by the signal handler's execute function when it
// terminates due to a received signal.
//
// SignalError has a design error that impacts comparison with errors.As.
// Callers should prefer using errors.Is(err, ErrSignal) to check for signal
// errors, and should only use errors.As in the rare case that they need to
// program against the specific os.Signal value.
type SignalError struct {
	Signal os.Signal
}

// Error implements the error interface.
//
// It was a design error to define this method on a value receiver rather than a
// pointer receiver. For compatibility reasons it won't be changed.
func (e SignalError) Error() string { _ = "STUB: not implemented"; return "" }

// Is addresses a design error in the SignalError type, so that errors.Is with
// ErrSignal will return true.
func (e SignalError) Is(err error) bool { _ = "STUB: not implemented"; return false }

// As fixes a design error in the SignalError type, so that errors.As with the
// literal `&SignalError{}` will return true.
func (e SignalError) As(target interface{}) bool { _ = "STUB: not implemented"; return false }

// ErrSignal is returned by SignalHandler when a signal triggers termination.
var ErrSignal = errors.New("signal error")
