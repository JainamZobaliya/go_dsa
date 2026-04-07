package main

import (
	"fmt"
	"strings"
)

// Aprroach-1
// func isSubsequence(s string, t string) bool {
//     if len(s) == 0 {
//         return true
//     }
//     if len(s) > len(t) {
//         return false
//     }

//     sp := 0

//     for tp := 0; tp < len(t); tp++ {
//         if s[sp] == t[tp] {
//             sp++
//         }
//         if sp == len(s) {
//             return true
//         }
//     }

//     return false
// }


// Approach-2:
func isSubsequence(s string, t string) bool {
    sp := 0

    for tp := 0; tp < len(t); tp++ {
        if sp < len(s) && s[sp] == t[tp] {
            sp++
        }
        if sp == len(s) {
            return true
        }
    }

    return sp == len(s)
}

func main() {
	type testCase struct {
		s string
		t string
	}

	testCases := []testCase{
		{"abc", "ahbgdc"},  
		{"axc", "ahbgdc"},  
		{"", "ahbgdc"},     
		{"abc", ""},        
		{"ace", "abcde"}, 
	}

	for i, tc := range testCases {
		result := isSubsequence(tc.s, tc.t)
		fmt.Printf("%d.\ts: \"%s\"\n", i+1, tc.s)
		fmt.Printf("\tt: \"%s\"\n", tc.t)
		fmt.Printf("\n\tOutput: %t\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}