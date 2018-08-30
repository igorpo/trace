package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New() should not be nil")
	} else {
		test := "Hello trace package!"
		tracer.Trace(test)
		if buf.String() != test+"\n" {
			t.Errorf("Trace() error: got=%s, want=%s", buf.String(), test+"\n")
		}
	}
}
