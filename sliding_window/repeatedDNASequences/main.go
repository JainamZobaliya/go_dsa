package main

import (
	"fmt"
	"strings"
)

func findRepeatedDnaSequences(s string) []string {
	seen := make(map[string]int)
	result := []string{}
	for i := 0; i <= len(s)-10; i++ {
		substring := s[i : i+10]
		seen[substring]++
		if seen[substring] == 2 {
			result = append(result, substring)
		}
	}
	return result
}

func main() {
	testCases := []string{
		"ACGTACGTACGTACGT",
		"AAAAAAAAAAAAAAAAAAAA",
		"ACGTACGTACACGTACGTAC",
		"GAGAGAGAGAGAGAGAGAGA",
		"ATCGATCGATCGATCGATCGATCG",
	}
	for idx, s := range testCases {
		result := findRepeatedDnaSequences(s)
		fmt.Printf("%d.\ts: \"%s\"\n", idx+1, s)
		quoted := make([]string, len(result))
		for i, r := range result {
			quoted[i] = fmt.Sprintf("\"%s\"", r)
		}
		fmt.Printf("\tResult: [%s]\n", strings.Join(quoted, ", "))
		fmt.Println(strings.Repeat("-", 100))
	}
}