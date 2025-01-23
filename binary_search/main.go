package main

import "fmt"

//func sort(arr){
//
//}

func binarySearch(arr []int, target int) int {
	//sortedArr := sort(arr)
	startIndex := 0
	endIndex := len(arr) - 1
	for {
		middleIndex := startIndex + (endIndex-startIndex)/2
		middleValue := arr[middleIndex]
		if target == middleValue {
			return middleIndex
		} else if target > middleValue {
			startIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 5, 22, 34, 67}
	target := 34
	index := binarySearch(arr, target)
	fmt.Printf("%d", index)
}
