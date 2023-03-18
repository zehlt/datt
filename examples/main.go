package main

import "github.com/zehlt/datt"

func main() {
	uf := datt.NewUnionFind(5)
	uf.Union(4, 0)
	uf.Union(3, 1)
	uf.Union(4, 2)
	uf.Union(3, 2)

	uf.Find(4, 2)
}
