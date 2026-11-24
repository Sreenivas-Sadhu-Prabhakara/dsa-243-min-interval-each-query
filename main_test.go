package main

import (
	"reflect"
	"testing"
)

func TestMinIntervalPerQuery(t *testing.T) {
	got := MinIntervalPerQuery([][2]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5})
	if !reflect.DeepEqual(got, []int{3, 3, 1, 4}) {
		t.Errorf("got %v, want [3 3 1 4]", got)
	}

	got = MinIntervalPerQuery([][2]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})
	if !reflect.DeepEqual(got, []int{2, -1, 4, 6}) {
		t.Errorf("got %v, want [2 -1 4 6]", got)
	}

	// Queries out of order must come back in input order.
	got = MinIntervalPerQuery([][2]int{{1, 10}}, []int{9, 1})
	if !reflect.DeepEqual(got, []int{10, 10}) {
		t.Errorf("order = %v", got)
	}

	got = MinIntervalPerQuery(nil, []int{5})
	if !reflect.DeepEqual(got, []int{-1}) {
		t.Errorf("no intervals = %v", got)
	}
}
