package main

import (
	"fmt"
	"strings"
)

// Approach-1
// func removeDuplicates(nums []int) int{
    
//     p1, p2 := 0,1

// 	for p1<p2 {
// 		if p2>=len(nums) {
// 			break
// 		}
// 		if nums[p2] == nums[p1] {
// 			if p2 == len(nums) { // Pop element at p2 i.e. last index
// 				nums = append(nums[:p2])
// 			} else { // Pop element at p2
// 				nums = append(nums[:p2], nums[p2+1:]...) 
// 			}
// 		} else {
// 			p1++
// 			p2++
// 		}
// 	}

//     return len(nums)
// }


// Approach-2:
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    p1 := 0

    for p2 := 1; p2 < len(nums); p2++ {
        if nums[p2] != nums[p1] {
            p1++
            nums[p1] = nums[p2]
        }
    }

    return p1 + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2, 2, 3},
		{-1, -1, 0, 0, 1, 1, 2},
		{5, 5, 5, 5},
		{1, 2, 3, 4},
		{0},
	}

	for idx, nums := range testCases {
		fmt.Printf("%d.\tnums:%v\n", idx+1, nums)

		// because function modifies in-place
		arr := append([]int(nil), nums...)

		k := removeDuplicates(arr)

		fmt.Printf("\n\tUnique Count (k): %d\n", k)
		fmt.Printf("\tArray After Removing Duplicates: %v\n", arr[:k])
		fmt.Println()
		fmt.Println(strings.Repeat("-", 100))
	}
}
