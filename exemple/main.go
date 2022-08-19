package main

import (
	"log"

	"github.com/zehlt/datt"
)

func main() {
	// arr := []int{2, 4, 7, 8, 1, 9, 11}
	arr := []int{31, 4, 15, 8, 25, 9, 11, 7, 1, 3, 5, 5, 2}

	log.Println(arr)
	datt.MergeSort(arr)
	log.Println(arr)
}

// f, err := os.Create("testastos.prof")
// if err != nil {
// 	log.Fatal("err", err)
// }
// pprof.StartCPUProfile(f)
// defer pprof.StopCPUProfile()

// l := datt.NewLinkedList[string]()

// for t := 0; t < 100; t++ {
// 	for i := 0; i < 100_000; i++ {
// 		l.PushHead("boa")
// 	}

// 	for i := 0; i < 100_000; i++ {
// 		l.Clear()
// 	}
// }
