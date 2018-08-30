package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface for an object that can aid in tracing execution of code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New creates a new tracer with an embedded io.Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
