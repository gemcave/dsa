package main

import "fmt"

func swap(array []int, a, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}

func shellSorting(array []int, length int) {
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ { // 5
			var j = i // 5
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap) // 5, 0
				j = j - gap
			}
		}
	}
}

func main() {
	var scores = []int{9, 6, 5, 8, 0, 7, 4, 3, 2, 1, 2}
	var length = len(scores)

	shellSorting(scores, length)

	fmt.Println(scores)
}
