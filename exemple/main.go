package main

import (
	"log"

	"github.com/zehlt/datt"
)

func main() {
	a, _ := datt.NewSetVar("henry", "marie", "john")
	b := datt.NewSet[string]()

	b.Append("tyler")
	b.Append("alice")
	b.Append("henry")

	c := a.Union(b) // doesn't mutate a or b, create brand new set
	log.Println(c)  // (tyler, henry, marie, john, alice)

	d := b.Intersection(a)
	log.Println(d) // henry

	log.Println(a.IsSubset(b)) // false
}
