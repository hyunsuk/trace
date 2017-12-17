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

// reference followed code
// https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go#L110
type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// New is create tracer and return
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off is create nilTracer and return
func Off() Tracer {
	return &nilTracer{}
}
