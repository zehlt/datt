package main

import (
	"log"

	"github.com/zehlt/datt"
)

func main() {
	b, _ := datt.NewBitset(8)
	b.Set(3, true)
	b.Set(7, true)
	b.Set(0, true)
	b.Set(5, true)
	b.Set(6, true)

	b2, _ := datt.NewBitset(8)
	b2.Set(3, true)
	b2.Set(0, true)
	b2.Set(6, true)
	b2.Set(1, true)

	log.Println(b)
	log.Println(b2)
	log.Println(b.Contain(b2))
}
