package search_insert

import "testing"

func check(t *testing.T, actual interface{}, expected interface{}, msg string) {
	if actual != expected {
		t.Fatalf("Fail on case %s, expected: %d, but got: %d\n", msg, expected, actual)
	}
}

func TestSearchInsert(t *testing.T) {
	check(t, SearchInsert([]int{1, 3, 5, 6}, 5), 2, "1")
	check(t, SearchInsert([]int{1, 3, 5, 6}, 2), 1, "2")
	check(t, SearchInsert([]int{1, 3, 5, 6}, 7), 4, "3")
	check(t, SearchInsert([]int{2, 3, 5, 6}, 1), 0, "4")
}
