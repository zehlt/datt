package datt

import "log"

// O(n²)
func InsertionSort(arr []int) []int {

	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}

	return arr
}

func merge(arr []int, p int, q int, r int) {
	nL := q - p + 1
	nR := r - q

	L := make([]int, nL)
	R := make([]int, nR)

	for i := 0; i < nL; i++ {
		L[i] = arr[p+i]
	}

	for j := 0; j < nR; j++ {
		R[j] = arr[q+j+1]
	}

	i := 0
	j := 0
	k := p

	for i < nL && j < nR {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < nL {
		arr[k] = L[i]
		i++
		k++
	}

	for j < nR {
		arr[k] = R[j]
		j++
		k++
	}
}

func mergeSortRec(arr []int, p int, r int) {
	if p >= r {
		return
	}

	q := (p + r) / 2
	mergeSortRec(arr, p, q)
	mergeSortRec(arr, q+1, r)
	log.Println("merge: ", p, q, r)
	merge(arr, p, q, r)
}

func MergeSort(arr []int) {
	p := 0
	r := len(arr) - 1
	mergeSortRec(arr, p, r)
}
