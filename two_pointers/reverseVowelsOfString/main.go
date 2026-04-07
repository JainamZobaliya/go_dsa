package main

func reverseVowels(s string) string {
	chars := []byte(s)
    var left int = 0
    var right int = len(chars) - 1
	// var VOWELS = []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	vowels := map[byte]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
    }

	// Approach-1
	// for ; left<=right; {
	// 	if !vowels[chars[left]]{
	// 		left++
	// 		continue
	// 	} else if !vowels[chars[right]]{
	// 		right--
	// 		continue
	// 	} else {
	// 		chars[left], chars[right] = chars[right], chars[left]
	// 		left++
	// 		right--
	// 	}
	// }

	// Approach-2
	for left < right {
		for left < right && !vowels[chars[left]] {
			left++
		}
		for left < right && !vowels[chars[right]] {
			right--
		}
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return string(chars)
}
