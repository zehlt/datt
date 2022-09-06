package datt

// // O(n²)
// func SortBubble[T constraints.Ordered](arr []T) []T {
// 	length := len(arr)

// 	for i := length - 1; i > 0; i-- {
// 		for y := 0; y < i; y++ {
// 			first := arr[y]
// 			second := arr[y+1]

// 			if first > second {
// 				arr[y+1] = first
// 				arr[y] = second
// 			}
// 		}
// 	}

// 	return arr
// }

// // O(n²)
// func SortSelection[T constraints.Ordered](arr []T) []T {
// 	length := len(arr)

// 	// while index is not last
// 	for i := 0; i < length; i++ {

// 		// find the smallest element
// 		lowest_value := arr[i]
// 		lowest_index := i
// 		for y := i + 1; y < length; y++ {
// 			if lowest_value > arr[y] {
// 				lowest_value = arr[y]
// 				lowest_index = y
// 			}
// 		}

// 		// swap with the index
// 		if i != lowest_index {
// 			previous_first := arr[i]
// 			arr[i] = lowest_value
// 			arr[lowest_index] = previous_first
// 		}
// 	}

// 	return arr
// }

// func SortInsertion[T constraints.Ordered](arr []T) []T {

// 	for current_index, current_value := range arr {
// 		previous_index := current_index - 1

// 		for previous_index >= 0 {
// 			if arr[previous_index] > current_value {
// 				arr[previous_index+1] = arr[previous_index]
// 			} else {
// 				break
// 			}
// 			previous_index--
// 		}
// 		arr[previous_index+1] = current_value
// 	}

// 	return arr
// }

// func partition[T constraints.Ordered](arr []T, leftPtr int, rightPtr int) int {
// 	pivotIndex := rightPtr
// 	pivotValue := arr[pivotIndex]
// 	rightPtr -= 1

// 	for {
// 		for arr[leftPtr] < pivotValue {
// 			leftPtr += 1

// 			if leftPtr > len(arr)-1 {
// 				break
// 			}
// 		}

// 		for arr[rightPtr] > pivotValue {
// 			rightPtr -= 1
// 			if rightPtr <= 0 {
// 				break
// 			}
// 		}

// 		if leftPtr >= rightPtr {
// 			break
// 		} else {
// 			temp := arr[leftPtr]
// 			arr[leftPtr] = arr[rightPtr]
// 			arr[rightPtr] = temp
// 			leftPtr += 1
// 		}
// 	}

// 	temp := arr[leftPtr]
// 	arr[leftPtr] = arr[pivotIndex]
// 	arr[pivotIndex] = temp

// 	return leftPtr
// }

// func sortQuickRec[T constraints.Ordered](arr []T, leftIndex int, rightIndex int) {
// 	if rightIndex-leftIndex <= 0 {
// 		return
// 	}

// 	pivotIndex := partition(arr, leftIndex, rightIndex)

// 	sortQuickRec(arr, leftIndex, pivotIndex-1)
// 	sortQuickRec(arr, pivotIndex+1, rightIndex)
// }

// func SortQuick[T constraints.Ordered](arr []T) []T {
// 	sortQuickRec(arr, 0, len(arr)-1)
// 	return arr
// }
