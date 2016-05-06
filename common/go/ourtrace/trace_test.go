package ourtrace

import (
	"testing"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
)

func TestNewContext(t *testing.T) {
	traceID := uint64(1)
	spanID := uint64(2)

	tr := New("foo", "bar")
	tr.SetTraceInfo(traceID, spanID)

	ctx := trace.NewContext(context.Background(), tr)
	tr, ok := trace.FromContext(ctx)
	if !ok {
		t.Fatal("context should contain our trace")
	}
	gotr := tr.(*Trace)

	if got, want := gotr.TraceID, traceID; got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
	if got, want := gotr.SpanID, spanID; got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
