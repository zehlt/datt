package main

import (
	"log"

	"github.com/zehlt/datt"
)

func main() {
	trie := datt.NewTrie()

	trie.Insert("car")
	trie.Insert("cartridge")
	trie.Insert("plane")
	trie.Insert("planer")

	ca := trie.AutoComplete("ca")
	log.Println(ca) // [r, rtridge]

	p := trie.AutoComplete("p")
	log.Println(p) // [lane, laner]

	doc := trie.AutoComplete("doc")
	log.Println(doc) // []
}
