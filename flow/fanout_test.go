package flow_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/micheam/go-steam/flow"
	"github.com/micheam/go-steam/sink"
	"github.com/micheam/go-steam/source"
)

func TestBroadcast(t *testing.T) {
	var (
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
		orig        = []int{1, 2, 3, 4, 5}
		src         = source.Stream(ctx, orig)
	)
	defer cancel()

	n := 5
	got := flow.Broadcast(ctx, n, src)
	if len(got) != n {
		t.Errorf("want %d, but got %d", n, len(got))
		t.FailNow()
	}

	coll := sink.Collect(flow.Merge(ctx, got))
	sort.IntSlice(coll).Sort()
	if diff := cmp.Diff(orig, coll); diff != "" {
		t.Errorf("Merged mismatch (-want, +got):%s\n", diff)
	}
}