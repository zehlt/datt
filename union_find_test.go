package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestUFNew(t *testing.T) {

	t.Run("on correct value must return non nil", func(t *testing.T) {
		uf := datt.NewUnionFind(10)

		AssertNotEqual(t, uf, nil)
	})

	t.Run("on negative value must return non nil", func(t *testing.T) {
		uf := datt.NewUnionFind(-31)

		AssertEqual(t, uf, nil)
	})
}

func TestUFClear(t *testing.T) {
	t.Run("clear remove all unions", func(t *testing.T) {
		uf := datt.NewUnionFind(10)

		uf.Union(1, 2)
		got := uf.Find(1, 2)
		AssertEqual(t, got, true)

		uf.Clear()
		got = uf.Find(1, 2)
		AssertEqual(t, got, false)
	})
}

func TestUFFind(t *testing.T) {
	t.Run("not union", func(t *testing.T) {
		uf := datt.NewUnionFind(10)

		got := uf.Find(5, 3)

		AssertEqual(t, got, false)
	})

	t.Run("out of range must return false", func(t *testing.T) {
		uf := datt.NewUnionFind(10)

		got := uf.Find(37, 3)

		AssertEqual(t, got, false)
	})

	t.Run("negative numbers must return false", func(t *testing.T) {
		uf := datt.NewUnionFind(10)

		got := uf.Find(-2, 3)

		AssertEqual(t, got, false)
	})

	t.Run("find after union must return true in both way", func(t *testing.T) {
		uf := datt.NewUnionFind(5)

		uf.Union(2, 1)
		got := uf.Find(2, 1)

		AssertEqual(t, got, true)
	})

	t.Run("union to itself", func(t *testing.T) {
		uf := datt.NewUnionFind(5)

		uf.Union(3, 3)
		got := uf.Find(3, 3)

		AssertEqual(t, got, true)
	})

	t.Run("union out of range", func(t *testing.T) {
		uf := datt.NewUnionFind(5)

		got := uf.Union(5, 3)

		AssertEqual(t, got, false)
	})
}

type findquery struct {
	query [2]int
	want  bool
}

func TestUnionFind(t *testing.T) {
	testCases := []struct {
		desc   string
		cap    int
		unions [][2]int
		finds  []findquery
	}{
		{
			desc: "",
			cap:  5,
			unions: [][2]int{
				{2, 1},
				{1, 2},
				{1, 1},
			},
			finds: []findquery{
				{[2]int{1, 2}, true},
				{[2]int{2, 1}, true},
				{[2]int{3, 4}, false},
				{[2]int{4, 4}, true},
			},
		},
		{
			desc: "",
			cap:  5,
			unions: [][2]int{
				{1, 2},
				{3, 4},
				{4, 3},
				{2, 3},
				{3, 2},
			},
			finds: []findquery{
				{[2]int{1, 2}, true},
				{[2]int{3, 4}, true},
				{[2]int{1, 4}, true},
				{[2]int{2, 4}, true},
				{[2]int{1, 3}, true},
				{[2]int{0, 3}, false},
				{[2]int{0, 1}, false},
			},
		},
		{
			desc: "test with weight b bigger",
			cap:  5,
			unions: [][2]int{
				{3, 4},
				{1, 2},
				{0, 3},
				{1, 3},
			},
			finds: []findquery{
				{[2]int{0, 1}, true},
				{[2]int{0, 3}, true},
				{[2]int{1, 2}, true},
				{[2]int{1, 4}, true},
				{[2]int{1, 3}, true},
				{[2]int{2, 4}, true},
				{[2]int{3, 4}, true},
			},
		},
		{
			desc: "test with weight a bigger",
			cap:  5,
			unions: [][2]int{
				{4, 0},
				{3, 1},
				{4, 2},
				{3, 2},
			},
			finds: []findquery{
				{[2]int{0, 1}, true},
				{[2]int{0, 3}, true},
				{[2]int{1, 2}, true},
				{[2]int{1, 4}, true},
				{[2]int{1, 3}, true},
				{[2]int{2, 4}, true},
				{[2]int{3, 4}, true},
			},
		},
		{
			desc: "all union",
			cap:  5,
			unions: [][2]int{
				{1, 2},
				{3, 4},
				{3, 2},
				{0, 2},
			},
			finds: []findquery{
				{[2]int{1, 2}, true},
				{[2]int{3, 4}, true},
				{[2]int{1, 4}, true},
				{[2]int{2, 4}, true},
				{[2]int{1, 3}, true},
				{[2]int{0, 3}, true},
				{[2]int{0, 1}, true},
			},
		},
		{
			desc: "with path compression",
			cap:  10,
			unions: [][2]int{
				{0, 1},
				{0, 2},
				{3, 4},
				{0, 3},
				{5, 6},
				{7, 8},
				{5, 3},
			},
			finds: []findquery{
				{[2]int{1, 2}, true},
				{[2]int{3, 4}, true},
				{[2]int{1, 4}, true},
				{[2]int{2, 4}, true},
				{[2]int{1, 3}, true},
				{[2]int{0, 3}, true},
				{[2]int{0, 1}, true},
				{[2]int{5, 6}, true},
				{[2]int{5, 3}, true},
				{[2]int{7, 8}, true},
				{[2]int{5, 8}, false},
				{[2]int{8, 9}, false},
				{[2]int{0, 9}, false},
				{[2]int{8, 3}, false},
				{[2]int{4, 8}, false},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			uf := datt.NewUnionFind(tC.cap)
			AssertNotEqual(t, uf, nil)

			for _, un := range tC.unions {
				uf.Union(un[0], un[1])
			}

			for _, fi := range tC.finds {
				got := uf.Find(fi.query[0], fi.query[1])

				if got != fi.want {
					t.Errorf("find(%d, %d) got %v want %v", fi.query[0], fi.query[1], got, fi.want)
				}
			}
		})
	}
}
