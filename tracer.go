package trace

import (
	"fmt"
	"io"
)

// Tracer is ...
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

// New is return tracer referance
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
