package main

import (
    "fmt"
	"strings"
)

func findMaxSlidingWindow(nums []int, w int) []int {
	result := make([]int, 0)
	prevMaxIndex := 0

	for i:=0; i<len(nums)-w+1; i++ {
		if i >= prevMaxIndex {
			prevMaxIndex = i
		}
		window := nums[prevMaxIndex:i+w]
		currMax := window[0]
		for j:=1; j<len(window); j++ {
			if (window[j]>=currMax) {
				prevMaxIndex = i+j
				currMax = window[j]
			}
		}
		result = append(result, currMax)
	}

	return result
}


func main() {
    windowSizes := []int{3, 3, 3, 3, 2, 4, 3, 2, 3, 6}
    numsList := [][]int{
        {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
        {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
        {10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
        {1, 5, 8, 10, 10, 10, 12, 14, 15, 19, 19, 19, 17, 14, 13, 12, 12, 12, 14, 18, 22, 26, 26, 26, 28, 29, 30},
        {10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67},
		{4, 5, 6, 1, 2, 3},
		{9, 5, 3, 1, 6, 3},
		{2, 4, 6, 8, 10, 12, 14, 16},
		{-1, -1, -2, -4, -6, -7},
		{4, 4, 4, 4, 4, 4},
	}

	for index, value := range numsList {
		fmt.Printf("%d.\tInput array:\t%s\n", index + 1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tWindow size:\t%d\n", windowSizes[index])
		output := findMaxSlidingWindow(value, windowSizes[index])
		fmt.Printf("\n\tMaximum in each sliding window:\t%s\n", strings.Replace(fmt.Sprint(output), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
