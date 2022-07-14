package main

import (
	"github.com/zehlt/datt"
)

func main() {
	l := datt.NewLinkedList[string]()
	l.PushHead("salut")
	l.PushHead("aoeee")

	l.PushTail("many")
	l.PushTail("toe")

	// l.PushTail("aaa")
	// l.Map(func(data string) {
	// 	fmt.Println(data)
	// })

	// fmt.Println(l.String())
}
